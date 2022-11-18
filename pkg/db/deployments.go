package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type CreateDeployment struct {
	Image string `json:"image" binding:"required"`
	Version string `json:"version" binding:"semver"`
}

func (d *DDBService) CreateDeployment() {
	// var item Deployment
	depl := CreateDeployment{Image: "my-image-repo/my-image", Version: "0.0.1"}

	av, err := dynamodbattribute.MarshalMap(depl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("Deployments"),
	}

	_, err = d.ddb.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	
	fmt.Println("Success!")
}