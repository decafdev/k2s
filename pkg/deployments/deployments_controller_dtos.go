package deployments

// DeploymentDTO struct
type DeploymentDTO struct {
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Environment map[string]string `json:"environment"`
}
