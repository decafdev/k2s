package sdk

import "github.com/techdecaf/k2s/v2/pkg/deployments"

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
