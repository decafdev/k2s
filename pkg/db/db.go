package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

type DDBService struct {
	ddb *dynamodb.DynamoDB
}

func NewDDB(config *config.ConfigService) (*DDBService, error) {
	awsConfig := &aws.Config{
		Region: aws.String(config.AWS_REGION),
		Endpoint: aws.String(config.DDB_URL),
	}

	sess := session.Must(session.NewSession(awsConfig))

	ddb := dynamodb.New(sess)

	ddbService := &DDBService{ddb: ddb}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("image"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("version"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("image"),
				KeyType: aws.String("HASH"),	
			},
			{
				AttributeName: aws.String("version"),
				KeyType: aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Deployments"),
	}

	_, err := ddbService.ddb.CreateTable(input)
	if err != nil && err.Error() != "ResourceInUseException: Cannot create preexisting table" {
		return ddbService, err
	}

	return ddbService, nil
}