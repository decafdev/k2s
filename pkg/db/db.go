package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

type DDBService struct {
	ddb *dynamodb.Client
}

func NewDDB(config *config.ConfigService) (*DDBService, error) {
	cfg, err := cfg.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	ddb := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL(config.DDB_URL)
	})

	ddbService := &DDBService{ddb: ddb}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("image"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("version"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("image"),
				KeyType: types.KeyTypeHash,	
			},
			{
				AttributeName: aws.String("version"),
				KeyType: types.KeyTypeRange,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		TableName: aws.String("Deployments"),
	}

	_, err = ddbService.ddb.CreateTable(context.Background(), input)
	target := &types.ResourceInUseException{}
	if err != nil && !errors.As(err, &target) {
		fmt.Println(err.Error())
		return ddbService, err
	}

	return ddbService, nil
}