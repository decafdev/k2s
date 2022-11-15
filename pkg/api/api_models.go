package api


type getHealthResponse struct {
	Name     string `json:"name" example:"my-app"`
	Version  string `json:"version" example:"1.0.1"`
	Hostname string `json:"hostname" example:"api.my-app.com"`
}

type createDeploymentRequest struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}

type createDeploymentResponse struct {
	Name        string            `json:"name" binding:"required" example:"my-service"`
	Image       string            `json:"image" binding:"required" example:"techdecaf/k2s"`
	Version     string            `json:"version" binding:"semver" example:"1.0.1"`
	Environment map[string]string `json:"environment"`
}
