package db

type CreateDeployment struct {
	Image string `json:"image" binding:"required"`
	Version string `json:"version" binding:"semver"`
}