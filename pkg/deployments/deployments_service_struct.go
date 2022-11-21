package deployments

import (
	"fmt"

	"github.com/techdecaf/k2s/v2/pkg/kube"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDeploymentService function description
func NewDeploymentService(k8s *kube.Service) DeploymentSrv {
	return &DeploymentService{k8s: k8s, labels: CommonLabels{
		CreatedBy: "k2s-operator",
	}}
}

type DeploymentSrv interface {
	CreateDeployment(spec *DeploymentDTO) error
	ListNamespaces() ([]v1.Namespace, error)
	ListDeployments() ([]DeploymentStatus, error)
}

// CommonLabels struct
type CommonLabels struct {
	CreatedBy string
}

// DeploymentService struct
type DeploymentService struct {
	labels CommonLabels
	// table *state.DeploymentsTable
	k8s *kube.Service
}

// CreateDeployment method
func (t *DeploymentService) CreateDeployment(spec *DeploymentDTO) error {
	application, err := kube.NewAPIApplication(&kube.APIOptions{
		Name:        spec.Name,
		Image:       spec.Image,
		Port:        80,
		Version:     spec.Version,
		Replicas:    3,
		MemoryLimit: 0,
		CPULimit:    0,
		Variables:   map[string]string{},
		Middlewares: []string{},
		CreatedBy:   t.labels.CreatedBy,
	})
	if err != nil {
		return err
	}

	return application.Apply(t.k8s)

	// return t.table.Create(spec)
}

// ListNamespaces method
func (t *DeploymentService) ListNamespaces() ([]v1.Namespace, error) {
	res, err := t.k8s.ListNamespaces(metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", "created-by", t.labels.CreatedBy),
	})
	return res.Items, err
}

// ListDeployments method
func (t *DeploymentService) ListDeployments() ([]DeploymentStatus, error) {
	deployments := []DeploymentStatus{}
	namespaces, err := t.ListNamespaces()
	if err != nil {
		return nil, err
	}

	for _, namespace := range namespaces {
		list, err := t.k8s.ListDeployments(namespace.Name, metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", "created-by", t.labels.CreatedBy),
		})
		if err != nil {
			return nil, err
		}

		for _, deployment := range list.Items {
			deployments = append(deployments, DeploymentStatus{
				Name:      deployment.Name,
				Namespace: namespace.Name,
				Image:     deployment.Spec.Template.Spec.Containers[0].Image,
				Version:   deployment.Labels["version"],
				Status:    deployment.Status,
			})
		}

	}
	return deployments, nil
}

// DeleteDeployment method
// func (t *DeploymentService) DeleteDeployment(key string) (err error) {
// 	return t.table.Delete(key)
// }

// OnModuleInit method
// func (t *DeploymentService) OnModuleInit() error {
// 	onCreate := t.table.OnChange().
// 		Filter(streams.Rx().KVOperationFilter(nats.KeyValuePut)).
// 		Map(t.table.Deserialize)

// 	onDelete := t.table.OnChange().
// 		Filter(streams.Rx().KVOperationFilter(nats.KeyValueDelete)).
// 		Map(t.table.DeserializeKey)

// 	go func() {
// 		for i := range onCreate.Observe() {
// 			d := i.V.(*dto.Deployment)
// 			t.log.Warnf("CREATE [id:%s] [name:%s] [version:%s]", d.Metadata.Id, d.Metadata.Name, d.Metadata.Version)
// 		}
// 	}()
// 	go func() {
// 		for i := range onDelete.Observe() {
// 			d := i.V.(*dto.Deployment)
// 			t.log.Warnf("DELETE [id:%s] [name:%s] [version:%s]", d.Metadata.Id, d.Metadata.Name, d.Metadata.Version)

// 		}
// 	}()
// 	return nil
// }
