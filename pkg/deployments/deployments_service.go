package deployments

import (
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/kube"
)

// DeploymentService struct
type DeploymentService struct {
	// table *state.DeploymentsTable
	kubeClient *kube.Client
	logger *logrus.Entry
}

// NewDeploymentService function description
func NewDeploymentService(kube *kube.Client, log *logrus.Entry) *DeploymentService {
	return &DeploymentService{kubeClient: kube, logger: log}
}

// CreateDeployment method
func (d *DeploymentService) CreateDeployment(depl *CreateDeploymentModel) error {
	application, err := NewAPIApplication(&APIOptions{
		Name:        depl.Name,
		Image:       depl.Image,
		Port:        80,
		Version:     depl.Version,
		Replicas:    3,
		MemoryLimit: 0,
		CPULimit:    0,
		Variables:   map[string]string{},
		Middlewares: []string{},
	})
	if err != nil {
		d.logger.Error(err)
		return err
	}

	application.Apply(d.kubeClient)
	return nil

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
