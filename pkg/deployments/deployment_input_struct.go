package deployments

// DeploymentDTO struct
type DeploymentDTO struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Port        string            `json:"port" binding:"required" example:"8080"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}
