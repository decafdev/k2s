package deployments

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/dto"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/state"
	"github.com/techdecaf/k2s/v2/pkg/streams"
)

// NewDeploymentService function description
func NewDeploymentService(table *state.DeploymentsTable, k8s *kube.Service, log *logrus.Entry) *DeploymentService {
	return &DeploymentService{k8s: k8s, log: log, table: table}
}

// DeploymentService struct
type DeploymentService struct {
	table *state.DeploymentsTable
	k8s   *kube.Service
	log   *logrus.Entry
}

// CreateDeployment method
func (t *DeploymentService) CreateDeployment(spec *state.DeploymentDTO) (*dto.Deployment, error) {
	return t.table.Create(spec)
}

// DeleteDeployment method
func (t *DeploymentService) DeleteDeployment(key string) (err error) {
	return t.table.Delete(key)
}

// OnModuleInit method
func (t *DeploymentService) OnModuleInit() error {
	onCreate := t.table.OnChange().
		Filter(streams.Rx().KVOperationFilter(nats.KeyValuePut)).
		Map(t.table.Deserialize)

	onDelete := t.table.OnChange().
		Filter(streams.Rx().KVOperationFilter(nats.KeyValueDelete)).
		Map(t.table.DeserializeKey)

	go func() {
		for i := range onCreate.Observe() {
			d := i.V.(*dto.Deployment)
			t.log.Infof("deploying [id:%s] [name:%s] [version:%s]", d.Metadata.Id, d.Metadata.Name, d.Metadata.Version)
		}
	}()
	go func() {
		for i := range onDelete.Observe() {
			d := i.V.(*dto.Deployment)
			t.log.Infof("deleting [name: %s] [version: %s]", d.Metadata.Name, d.Metadata.Version)
		}
	}()
	return nil
}
