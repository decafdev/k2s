package state

import v1 "k8s.io/api/apps/v1"

// DeploymentDTO struct
type DeploymentDTO struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}

// DeploymentDTO struct
type DeploymentStatus struct {
	Name      string              `json:"name"  example:"my-service"`
	Namespace string              `json:"namespace"  example:"whoami-v1"`
	Image     string              `json:"image"  example:"techdecaf/k2s"`
	Version   string              `json:"version" example:"1.0.1"`
	Status    v1.DeploymentStatus `json:"status"`
}
