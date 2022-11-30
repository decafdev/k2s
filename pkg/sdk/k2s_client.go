package sdk

import (
	"fmt"

	"github.com/techdecaf/k2s/v2/pkg/deployments"
)

// Client struct
type Client struct {
	sdk *Options
}

func NewClient(baseURI string) *Client {
	return &Client{sdk: &Options{
		BaseURI: baseURI,
	}}
}

// ListDeployments method
func (t *Client) ListDeployments() (res *[]deployments.DeploymentStatus, err error) {
	_, err = t.sdk.GetI("/deployments", &res)
	return res, err
}

// GetDeployment method
func (t *Client) GetDeployment(id string) (res *deployments.DeploymentStatus, err error) {
	_, err = t.sdk.GetI(fmt.Sprintf("/deployments/%s", id), &res)
	return res, err
}

// CreateDeployment method
func (t *Client) CreateDeployment(body deployments.DeploymentDTO) (res *deployments.DeploymentStatus, err error) {
	_, err = t.sdk.PostI("/deployments", body, &res)
	return res, err
}
