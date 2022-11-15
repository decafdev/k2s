package deployments

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/util"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	netV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// APIResources struct
type APIOptions struct {
	Name        string `validate:"required" mod:"lcase"`
	Image       string `validate:"required" mod:"lcase"`
	Port        int32  `validate:"required"`
	Version     string `validate:"required,semver"`
	Replicas    int32  `mod:"default=1"`
	MemoryLimit int64  `mod:"default=128"`
	CPULimit    int64  `mod:"default=500"`
	Variables   map[string]string
	Middlewares []string `validate:"unique,dive,required,endswith=@file" mod:"lcase"`
}

func (t *APIOptions) Validate() (*APIOptions, error) {
	if err := modifiers.New().Struct(context.Background(), t); err != nil {
		return t, err
	}
	return t, validator.New().Struct(t)
}

// APIResources struct
type APIResources struct {
	Namespace  *coreV1.Namespace
	Ingress    *netV1.Ingress
	Service    *coreV1.Service
	Deployment *appsV1.Deployment
	Secret     *coreV1.Secret
}

// Apply method
func (r *APIResources) Apply(client *kube.Client) {
	client.ApplyNamespace(r.Namespace)
	client.ApplyService(r.Service.ObjectMeta.Namespace, r.Service)
	client.ApplySecret(r.Secret.ObjectMeta.Namespace, r.Secret)
	client.ApplyDeployment(r.Deployment.ObjectMeta.Namespace, r.Deployment)
	client.ApplyIngress(r.Deployment.ObjectMeta.Namespace, r.Ingress)
}

// Application function description
func NewAPIApplication(o *APIOptions) (*APIResources, error) {
	o, err := o.Validate()
	if err != nil {
		return &APIResources{}, err
	}

	version, err := semver.NewVersion(o.Version)
	if err != nil {
		return &APIResources{}, err
	}

	metadata := metaV1.ObjectMeta{
		Name:      o.Name,
		Namespace: fmt.Sprintf("%s-v%v", o.Name, version.Major()),
		Labels: map[string]string{
			"app.kubernetes.io/name": o.Name,
		},
	}

	namespace := &coreV1.Namespace{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Namespace",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:   metadata.Namespace,
			Labels: metadata.Labels,
		},
	}

	secret := &coreV1.Secret{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      fmt.Sprintf("%s.%s", o.Name, o.Version),
			Namespace: namespace.ObjectMeta.Name,
			Labels:    metadata.Labels,
		},
		StringData: o.Variables,
	}

	return &APIResources{
		Namespace: namespace,
		Secret:    secret,
		Ingress: &netV1.Ingress{
			TypeMeta: metaV1.TypeMeta{APIVersion: "networking.k8s.io/v1", Kind: "Ingress"},
			ObjectMeta: metaV1.ObjectMeta{
				Name:      metadata.Name,
				Namespace: namespace.ObjectMeta.Name,
				Labels:    metadata.Labels,
				Annotations: map[string]string{
					"traefik.ingress.kubernetes.io/router.middlewares": strings.Join(o.Middlewares, ","),
				},
			},
			Spec: netV1.IngressSpec{
				Rules: []netV1.IngressRule{
					{
						IngressRuleValue: netV1.IngressRuleValue{
							HTTP: &netV1.HTTPIngressRuleValue{
								Paths: []netV1.HTTPIngressPath{
									{
										Path:     fmt.Sprintf("/%s/v%v/", o.Name, version.Major()),
										PathType: util.PathType(netV1.PathTypePrefix),
										Backend: netV1.IngressBackend{
											Service: &netV1.IngressServiceBackend{
												Name: o.Name,
												Port: netV1.ServiceBackendPort{Number: o.Port},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Service: &coreV1.Service{
			TypeMeta: metaV1.TypeMeta{
				APIVersion: "v1",
				Kind:       "Service",
			},
			ObjectMeta: metadata,
			Spec: coreV1.ServiceSpec{
				Type: "ClusterIP",
				Ports: []coreV1.ServicePort{
					{Name: "http", Protocol: "TCP", Port: o.Port},
				},
				Selector: map[string]string{
					"app.kubernetes.io/name": metadata.Name,
				},
			},
		},
		Deployment: &appsV1.Deployment{
			TypeMeta: metaV1.TypeMeta{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			ObjectMeta: metadata,
			Spec: appsV1.DeploymentSpec{
				Replicas: util.Int32ToPtr(o.Replicas),
				Selector: &metaV1.LabelSelector{
					MatchLabels: map[string]string{
						"app.kubernetes.io/name": metadata.Name,
					},
				},
				Strategy: appsV1.DeploymentStrategy{
					Type: "RollingUpdate",
					RollingUpdate: &appsV1.RollingUpdateDeployment{
						MaxUnavailable: &intstr.IntOrString{
							IntVal: 0,
						},
						MaxSurge: &intstr.IntOrString{
							IntVal: 1,
						},
					},
				},
				Template: coreV1.PodTemplateSpec{
					ObjectMeta: metadata,
					Spec: coreV1.PodSpec{
						Containers: []coreV1.Container{
							{
								Name:  metadata.Name,
								Image: o.Image,
								Ports: []coreV1.ContainerPort{
									{Name: "http", ContainerPort: 80, Protocol: "TCP"},
								},
								EnvFrom: []coreV1.EnvFromSource{
									{
										SecretRef: &coreV1.SecretEnvSource{
											LocalObjectReference: coreV1.LocalObjectReference{
												Name: secret.ObjectMeta.Name,
											},
										},
									},
								},
								Resources: coreV1.ResourceRequirements{
									Limits: coreV1.ResourceList{
										coreV1.ResourceMemory: *resource.NewScaledQuantity(o.MemoryLimit, resource.Mega),
										coreV1.ResourceCPU:    *resource.NewScaledQuantity(o.CPULimit, resource.Milli),
									},
									Requests: coreV1.ResourceList{
										coreV1.ResourceMemory: *resource.NewScaledQuantity(o.MemoryLimit/8, resource.Mega),
										coreV1.ResourceCPU:    *resource.NewScaledQuantity(o.CPULimit/10, resource.Milli),
									},
								},
								// On failure the pod is killed restarted.
								LivenessProbe: &coreV1.Probe{
									SuccessThreshold: 1,
									TimeoutSeconds:   1,
									FailureThreshold: 3,
									PeriodSeconds:    30,
									ProbeHandler: coreV1.ProbeHandler{
										HTTPGet: &coreV1.HTTPGetAction{
											Path:   "/healthz",
											Scheme: "HTTP",
											Port: intstr.IntOrString{
												Type:   intstr.Int,
												IntVal: o.Port,
											},
										},
									},
								},
								// On failure the pod is taken out of service and no traffic routed
								ReadinessProbe: &coreV1.Probe{
									SuccessThreshold: 1,
									TimeoutSeconds:   1,
									FailureThreshold: 2,
									PeriodSeconds:    30,
									ProbeHandler: coreV1.ProbeHandler{
										HTTPGet: &coreV1.HTTPGetAction{
											Path:   "/healthz",
											Scheme: "HTTP",
											Port: intstr.IntOrString{
												Type:   intstr.Int,
												IntVal: o.Port,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, err
}

// // ToYAML method
// func (t *APIResources) ToYAML() ([]byte, error) {
// 	return tx.ResourcesToYAML([]runtime.Object{
// 		t.Namespace,
// 		t.Secret,
// 		t.Ingress,
// 		t.Service,
// 		t.Deployment,
// 	})
// }