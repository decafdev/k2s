package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *DDBService) CreateDeployment(arg CreateDeployment) error {
	item, err := attributevalue.MarshalMap(arg)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Deployments"),
	}

	_, err = d.ddb.PutItem(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (d *DDBService) GetDeployment(arg ReadDeployment) (Deployment, error) {
	var depl Deployment

	key, err := attributevalue.MarshalMap(arg)
	if err != nil {
		return depl, err
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("Deployments"),
	}

	result, err := d.ddb.GetItem(context.Background(), input)
	if err != nil {
		return depl, err
	}

	err = attributevalue.UnmarshalMap(result.Item, &depl)
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

	output, err := d.ddb.Scan(context.Background(), input)
	if err != nil {
		return depls, err
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &depls)
	if err != nil {
		return depls, err
	}

	return depls, nil
}
