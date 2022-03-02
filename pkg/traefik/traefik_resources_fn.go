package traefik

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var tx = &kube.Transformers{}

// Resource struct
type TraefikResources struct {
	ServiceAccount     *coreV1.ServiceAccount
	ClusterRole        *rbacV1.ClusterRole
	ClusterRoleBinding *rbacV1.ClusterRoleBinding
	ConfigMap          *coreV1.ConfigMap
	Deployment         *appsV1.Deployment
	Service            *coreV1.Service
}

// NewTraefikResources function description
func NewTraefikResources(o *ResourceOptions, cfg *TraefikConfig) (*TraefikResources, error) {
	// var ymlcfg string

	metadata := metaV1.ObjectMeta{
		Name:      o.Name,
		Namespace: o.Namespace,
		Labels: map[string]string{
			"app.kubernetes.io/name": o.Name,
		},
	}
	configuration, err := cfg.ToYAML()
	if err != nil {
		return &TraefikResources{}, errors.Wrap(err, "failed to generate traefik configuration yaml")
	}

	return &TraefikResources{
		ConfigMap: &coreV1.ConfigMap{
			ObjectMeta: metadata,
			Data: map[string]string{
				"traefik-middlewares.yaml": string(configuration),
			},
		},
		ServiceAccount: &coreV1.ServiceAccount{
			ObjectMeta: metaV1.ObjectMeta{Name: o.Name, Namespace: o.Namespace},
		},
		ClusterRole: &rbacV1.ClusterRole{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      o.Name,
				Namespace: o.Namespace,
			},
			Rules: []rbacV1.PolicyRule{
				{
					APIGroups: []string{"*"},
					Resources: []string{"*"},
					Verbs:     []string{"*"},
					// ResourceNames: []string{"*"},
					// NonResourceURLs: []string{""},
				},
			},
		},
		ClusterRoleBinding: &rbacV1.ClusterRoleBinding{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      o.Name,
				Namespace: o.Namespace,
			},
			RoleRef: rbacV1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     o.Name,
			},
			Subjects: []rbacV1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      o.Name,
					Namespace: o.Namespace,
				},
			},
		},
		Deployment: &appsV1.Deployment{
			ObjectMeta: metadata,
			Spec: appsV1.DeploymentSpec{
				Replicas: tx.Int32ToPtr(o.Replicas),
				Strategy: appsV1.DeploymentStrategy{
					RollingUpdate: &appsV1.RollingUpdateDeployment{
						MaxSurge:       tx.IntOrStringI(1),
						MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 0},
					},
				},
				Selector: &metaV1.LabelSelector{
					MatchLabels: metadata.Labels,
				},
				Template: coreV1.PodTemplateSpec{
					ObjectMeta: metadata,
					Spec: coreV1.PodSpec{
						ServiceAccountName: o.Name,
						Volumes: []coreV1.Volume{
							{
								Name: o.Name,
								VolumeSource: coreV1.VolumeSource{
									ConfigMap: &coreV1.ConfigMapVolumeSource{
										LocalObjectReference: coreV1.LocalObjectReference{
											Name: o.Name,
										},
										// Items: []coreV1.KeyToPath{
										// 	{
										// 		Key:  "",
										// 		Path: "",
										// 		Mode: new(int32),
										// 	},
										// },
										// DefaultMode:          tx.Int32ToPtr(644),
										// Optional:             new(bool),
									},
								},
							},
						},
						// TerminationGracePeriodSeconds: 60,
						Containers: []coreV1.Container{
							{
								Image: fmt.Sprintf("traefik:v%s", o.Version),
								Name:  o.Name,
								Args: []string{
									"--ping=true",
									"--ping.entrypoint=http",
									"--global.checknewversion=true",
									"--api.debug=true",
									"--api.insecure=true",
									"--api.dashboard=true",
									"--accesslog",
									"--entryPoints.traefik.address=:8080",
									"--entrypoints.http.address=:8000",
									"--entrypoints.http.proxyprotocol",
									"--entrypoints.http.proxyprotocol.insecure",
									"--entrypoints.http.forwardedheaders.insecure",
									"--entrypoints.https.address=:4443",
									"--providers.kubernetesingress=true",
									"--providers.file.directory=/mounted",
									"--providers.file.watch=true",
									// "--providers.kubernetescrd",
									// "--metrics.datadog=true",
									// "--metrics.datadog.address=127.0.0.1:8125",
									// "--metrics.datadog.addEntryPointsLabels=true",
									// "--metrics.datadog.addServicesLabels=true",
									// "--metrics.datadog.pushInterval=10s",
									// // "--tracing.datadog=true",
									// "--tracing.datadog.localAgentHostPort=127.0.0.1:8126",
									// "--tracing.datadog.globalTag=sample",
									// "--tracing.datadog.prioritySampling=true",
									"--log.format=json",
									"--log.level=INFO",
								},
								Ports: []coreV1.ContainerPort{
									{Name: "http", ContainerPort: 8000, HostPort: o.HostHTTPPort},
									{Name: "https", ContainerPort: 4443},
									{Name: "dashboard", ContainerPort: 8080},
								},
								VolumeMounts: []coreV1.VolumeMount{
									{
										Name:      o.Name,
										ReadOnly:  true,
										MountPath: "/mounted",
									},
								},
								Resources: coreV1.ResourceRequirements{
									Limits: coreV1.ResourceList{
										coreV1.ResourceMemory: *resource.NewScaledQuantity(256, resource.Mega),
										coreV1.ResourceCPU:    *resource.NewScaledQuantity(1000, resource.Milli),
									},
									Requests: coreV1.ResourceList{
										coreV1.ResourceMemory: *resource.NewScaledQuantity(16, resource.Mega),
										coreV1.ResourceCPU:    *resource.NewScaledQuantity(100, resource.Milli),
									},
								},
								SecurityContext: &coreV1.SecurityContext{
									Capabilities: &coreV1.Capabilities{
										Drop: []coreV1.Capability{"ALL"},
										Add:  []coreV1.Capability{"NET_BIND_SERVICE"},
									},
								},
								ReadinessProbe: &coreV1.Probe{
									FailureThreshold: 1,
									ProbeHandler: coreV1.ProbeHandler{
										HTTPGet: &coreV1.HTTPGetAction{
											Path:   "/ping",
											Scheme: "HTTP",
											Port: intstr.IntOrString{
												Type:   intstr.Int,
												IntVal: 8000,
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
			ObjectMeta: metaV1.ObjectMeta{
				Name:      o.Name,
				Namespace: o.Namespace,
			},
			Spec: coreV1.ServiceSpec{
				Type: "NodePort",
				Selector: map[string]string{
					"app.kubernetes.io/name": o.Name,
				},
				Ports: []coreV1.ServicePort{
					{
						Name:       "dashboard",
						Port:       8080,
						TargetPort: *tx.IntOrStringI(8080),
						NodePort:   o.HostDashboardPort,
						Protocol:   "TCP",
					},
					{
						Name:       "http",
						Port:       80,
						TargetPort: *tx.IntOrStringI(8000),
						NodePort:   o.HostHTTPPort,
						Protocol:   "TCP",
					},
					{
						Name:       "https",
						Port:       443,
						TargetPort: *tx.IntOrStringI(8443),
						NodePort:   o.HostHTTPSPort,
						Protocol:   "TCP",
					},
				},
			},
		},
	}, nil
}
