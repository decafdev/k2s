package db

type Deployment struct {
	Name string `json:"name" binding:"required"`
	Version string `json:"version" binding:"semver"`
}

type CreateDeployment struct {
	Name string `json:"name" binding:"required" dynamodbav:"name"`
	Version string `json:"version" binding:"semver" dynamodbav:"version"`
}

type ReadDeployment struct {
	Name string `json:"name" binding:"required" dynamodbav:"name"`
	Version string `json:"version" binding:"semver" dynamodbav:"version"`
}