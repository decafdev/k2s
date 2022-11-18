package db

type Deployment struct {
	Image string `json:"image" binding:"required"`
	Version string `json:"version" binding:"semver"`
}