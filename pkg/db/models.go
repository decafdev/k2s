package db

type Deployment struct {
	Image string `json:"image" binding:"required"`
	Version string `json:"version" binding:"semver"`
}

type CreateDeployment struct {
	Image string `json:"image" binding:"required" dynamodbav:"image"`
	Version string `json:"version" binding:"semver" dynamodbav:"version"`
}

type ReadDeployment struct {
	Image string `json:"image" binding:"required" dynamodbav:"image"`
	Version string `json:"version" binding:"semver" dynamodbav:"version"`
}