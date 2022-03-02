package healthz

// HealthDTO struct
type HealthDTO struct {
	Name     string `json:"name" example:"my-app"`
	Version  string `json:"version" example:"1.0.1"`
	Hostname string `json:"hostname" example:"api.my-app.com"`
}
