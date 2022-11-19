package deployments

// DeploymentDTO struct
type createDeploymentRequest struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"required,semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}

type readDeploymentRequest struct {
	Name    string `uri:"name" binding:"required"`
	Version string `uri:"version" binding:"required,semver"`
}

type readDeploymentResponse struct {
	Name    string `json:"name" binding:"required" example:"techdecaf/k2s"`
	Version string `json:"version" binding:"required,semver" example:"1.0.1"`
}
