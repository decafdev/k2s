package state

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/reactivex/rxgo/v2"
	"github.com/techdecaf/k2s/v2/pkg/dto"
	"github.com/techdecaf/k2s/v2/pkg/streams"
	"google.golang.org/protobuf/proto"
)

// encodeKey function description
func encodeKey(keys ...string) string {
	// hash := md5.Sum([]byte(strings.Join(keys, ":")))
	return hex.EncodeToString([]byte(strings.Join(keys, ":")))
}

// decodeKey function description
func decodeKey(key string) ([]string, error) {
	b, err := hex.DecodeString(key)
	return strings.Split(string(b), ":"), err
}

// DeploymentDTO struct
type DeploymentDTO struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}

// DeploymentsTable struct
type DeploymentsTable struct {
	stream   *streams.Client
	table    nats.KeyValue
	onChange rxgo.Observable
}

// NewDeploymentsTable function description
func NewDeploymentsTable(stream *streams.Client) (*DeploymentsTable, error) {
	table, err := stream.Table("deployments")
	if err != nil {
		return &DeploymentsTable{}, err
	}

	return &DeploymentsTable{
		stream:   stream,
		table:    table,
		onChange: stream.OnTableEvent(table),
	}, nil
}

// Create method
func (t *DeploymentsTable) Create(item *DeploymentDTO) (spec *dto.Deployment, err error) {
	// generte key
	key := encodeKey(item.Name, item.Version)
	// convert to protobuf
	deployment := &dto.Deployment{
		Metadata: &dto.DeploymentMetadata{
			Id:      key,
			Type:    dto.DeploymentType_DEPLOYMENT_TYPE_API,
			Name:    item.Name,
			Version: item.Version,
		},
		Environment: item.Environment,
	}

	bytes, err := proto.Marshal(deployment)
	if err != nil {
		return spec, err
	}

	// call table create
	_, err = t.table.Create(key, bytes)
	return spec, err
}

// Get method
func (t *DeploymentsTable) Get(key string) (deployment *dto.Deployment, err error) {
	item, err := t.table.Get(key)
	if err != nil {
		return deployment, err
	}

	if err := proto.Unmarshal(item.Value(), deployment); err != nil {
		return deployment, err
	}

	return deployment, err
}

// Delete method
func (t *DeploymentsTable) Delete(key string) error {
	return t.table.Delete(key)
}

// OnChange method
func (t *DeploymentsTable) OnChange() rxgo.Observable {
	return t.onChange
}

// Deserialize method
func (t *DeploymentsTable) Deserialize(_ context.Context, item interface{}) (interface{}, error) {
	var deployment dto.Deployment
	i := item.(nats.KeyValueEntry)

	if err := proto.Unmarshal(i.Value(), &deployment); err != nil {
		return &deployment, err
	}
	return &deployment, nil
}

// DeserializeKey method
func (t *DeploymentsTable) DeserializeKey(_ context.Context, item interface{}) (interface{}, error) {
	i := item.(nats.KeyValueEntry)
	keys, err := decodeKey(i.Key())
	if err != nil {
		return &dto.Deployment{}, err
	}
	deployment := &dto.Deployment{
		Metadata: &dto.DeploymentMetadata{
			Id:      i.Key(),
			Type:    0,
			Name:    keys[0],
			Version: keys[1],
		},
	}
	return deployment, nil
}
