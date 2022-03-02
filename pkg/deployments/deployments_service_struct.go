package deployments

import (
	"github.com/nats-io/nats.go"
)

// NewDeploymentService function description
func NewDeploymentService(table nats.KeyValue) *DeploymentService {
	return &DeploymentService{
		table: table,
	}
}

// DeploymentService struct
type DeploymentService struct {
	table nats.KeyValue
}

// Deploy method
func (t *DeploymentService) Deploy() error {
	return nil
}
