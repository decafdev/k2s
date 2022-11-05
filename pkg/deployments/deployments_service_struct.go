package deployments

import (
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/state"
)

// NewDeploymentService function description
func NewDeploymentService(k8s *kube.Service, log *logrus.Entry) *DeploymentService {
	return &DeploymentService{k8s: k8s}
}

// DeploymentService struct
type DeploymentService struct {
	// table *state.DeploymentsTable
	k8s *kube.Service
	log *logrus.Entry
}

// CreateDeployment method
func (t *DeploymentService) CreateDeployment(spec *state.DeploymentDTO) error {
	apiResrouce, err := kube.NewAPIApplication(&kube.APIOptions{
		Name:        spec.Name,
		Image:       spec.Image,
		Port:        80,
		Version:     spec.Version,
		Replicas:    3,
		MemoryLimit: 0,
		CPULimit:    0,
		Variables:   map[string]string{},
		Middlewares: []string{},
	})
	if err != nil {
		t.log.Error(err)
	}

	return apiResrouce.Apply(t.k8s)

	// return t.table.Create(spec)
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
