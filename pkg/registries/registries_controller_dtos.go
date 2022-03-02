package registries

// CreateRegistryDTO struct
type CreateRegistryDTO struct {
	Name     string `json:"name" binding:"required" example:"docker-hub"`
	Registry string `json:"registry" binding:"required,url" example:"https://registry.docker.io"`
	Username string `json:"username" binding:"required" example:"my-user"`
	Password string `json:"password" binding:"required" example:"my-password"`
}

// PrivateRegistryDTO struct
type PrivateRegistryDTO struct {
	Name      string            `json:"name" example:"docker-hub"`
	Namespace string            `json:"namespace" example:"k2s-operator"`
	Labels    map[string]string `json:"labels"`
}
