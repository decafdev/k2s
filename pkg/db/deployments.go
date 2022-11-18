package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (d *DDBService) CreateDeployment(arg CreateDeployment) error {	
	item, err := dynamodbattribute.MarshalMap(arg)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String("Deployments"),
	}

	_, err = d.ddb.PutItem(input)
	if err != nil {
		return err
	}
		
	return nil
}

func (d *DDBService) GetDeployment(arg ReadDeployment) (Deployment, error) {
	var depl Deployment

	key, err := dynamodbattribute.MarshalMap(arg)
	if err != nil {
		return depl, err
	}

	input := &dynamodb.GetItemInput{
		Key: key,
		TableName: aws.String("Deployments"),
	}

	result, err := d.ddb.GetItem(input)
	if err != nil {
		return depl, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &depl)
	if err != nil {
		return depl, err
	}

	return depl, nil
}

func (d *DDBService) ListDeployments() ([]Deployment, error) {
	var depls []Deployment

	input := &dynamodb.ScanInput{
		TableName: aws.String("Deployments"),
	}

	output, err := d.ddb.Scan(input)
	if err != nil {
		return depls, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &depls)
	if err != nil {
		return depls, err
	}

	return depls, nil
}