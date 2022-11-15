package more_kube

// CreateRegistryDTO struct
type CreateRegistryDTO struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Registry  string            `json:"registry"`
	Username  string            `json:"username"`
	Password  string            `json:"password"`
	Labels    map[string]string `json:"labels"`
}

// PrivateRegistryDTO struct
type PrivateRegistryDTO struct {
	Name      string            `json:"name" example:"docker-hub"`
	Namespace string            `json:"namespace" example:"k2s-operator"`
	Labels    map[string]string `json:"labels"`
}
