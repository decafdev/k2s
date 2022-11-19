package deployments

import (
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/db"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/state"
)

// NewDeploymentService function description
func NewDeploymentService(k8s *kube.Service, log *logrus.Entry, ddb *db.DDBService) *DeploymentService {
	return &DeploymentService{k8s: k8s, log: log, ddb: ddb}
}

// DeploymentService struct
type DeploymentService struct {
	// table *state.DeploymentsTable
	k8s *kube.Service
	log *logrus.Entry
	ddb *db.DDBService
}

// CreateDeployment method
func (t *DeploymentService) CreateDeployment(spec *state.DeploymentDTO) error {
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
	})
	if err != nil {
		t.log.Error("failed to create new api app", err)
		return err
	}

	err = application.Apply(t.k8s)
	if err != nil {
		t.log.Error("failed to apply changes to cluster", err)
	}

	ddbItem := db.CreateDeployment{
		Name:    spec.Name,
		Version: spec.Version,
	}

	err = t.ddb.CreateDeployment(ddbItem)
	if err != nil {
		t.log.Error("failed to put deployment to ddb", err)
		return err
	}

	return nil

	// return t.table.Create(spec)
}

func (t *DeploymentService) GetDeployment(spec *state.DeploymentDTO) (*readDeploymentResponse, error) {
	ddbItem := db.ReadDeployment{
		Name:    spec.Name,
		Version: spec.Version,
	}

	depl, err := t.ddb.GetDeployment(ddbItem)
	if err != nil {
		t.log.Error("failed to get from ddb", err)
		return nil, err
	}

	resp := &readDeploymentResponse{
		Name:    depl.Name,
		Version: depl.Version,
	}

	return resp, nil
}

func (t *DeploymentService) ListDeployments() ([]readDeploymentResponse, error) {
	depls, err := t.ddb.ListDeployments()
	if err != nil {
		t.log.Error("failed to get from ddb", err)
		return nil, err
	}

	var resp []readDeploymentResponse
	for _, depl := range depls {
		item := readDeploymentResponse{
			Name:    depl.Name,
			Version: depl.Version,
		}
		resp = append(resp, item)
	}

	return resp, nil
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
