package deployments

import v1 "k8s.io/api/apps/v1"

// DeploymentStatus struct
type DeploymentStatus struct {
	Id        string              `json:"id"  example:"asdfasdf"`
	Name      string              `json:"name"  example:"my-service"`
	Namespace string              `json:"namespace"  example:"whoami-v1"`
	Image     string              `json:"image"  example:"techdecaf/k2s"`
	Version   string              `json:"version" example:"1.0.1"`
	Status    v1.DeploymentStatus `json:"status"`
}
