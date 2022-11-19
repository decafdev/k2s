package state

// DeploymentDTO struct
type DeploymentDTO struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"required,semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}
