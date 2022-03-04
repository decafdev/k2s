package deployments

// DeploymentDTO struct
type DeploymentRequest struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}
type DeploymentResponse struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}
