package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/require"
	"github.com/techdecaf/k2s/v2/pkg/util"
)

var create1 = CreateDeployment{
	Name: "image",
	Version: "0.0.0",
}

var create2 = CreateDeployment{
	Name: "my-image",
	Version: "0.99.0",
}

var read1 = ReadDeployment{
	Name: "image",
	Version: "0.0.0",
}

var read2 = ReadDeployment{
	Name: "my-image",
	Version: "0.99.0",
}

var create3 = CreateDeployment{}

var create4 = CreateDeployment{
	Version: "monotonic",
}

var read3 = ReadDeployment{
	Name: "my-image",
}

var read4 = ReadDeployment{
	Name: "doodle",
	Version: "0.1.0",
}

func TestCreateDeployment(t *testing.T) {
	type given struct {
		arg CreateDeployment
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when a deployment is properly specified"] = given{
		arg: create1,
	}

	testsNoErrors["when the user requests to deploy again"] = given{
		arg: create1,
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.NoError(t, err)
		})
	}

	testsErrors := make(map[string]given)

	testsErrors["when no deployment info is specified"] = given{
		arg: create3,
	}

	testsErrors["when a deployment is improperly specified"] = given{
		arg: create4,
	}

	for when, given := range testsErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.Error(t, err)
		})
	}

	cleanup()
}

func TestGetDeployment(t *testing.T) {	
	type given struct {
		createArg CreateDeployment
		readArg ReadDeployment
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when a deployment is in the db"] = given{
		createArg: create1,
		readArg: read1,
	}

	testsNoErrors["when another deployment is in the db"] = given{
		createArg: create2,
		readArg: read2,
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			createDeploymentNoErrors(t, given.createArg)
			depl, err := ddbService.GetDeployment(given.readArg)
			require.NoError(t, err)
			require.NotEmpty(t, depl)
			require.Equal(t, depl.Name, given.createArg.Name)
			require.Equal(t, depl.Version, given.createArg.Version)
			cleanup()
		})
	}

	testsNoItem := make(map[string]given)

	testsNoItem["when there is no matching item"] = given{
		readArg: read4,
	}

	for when, given := range testsNoItem {
		t.Run(when, func(t *testing.T) {
			depl, err := ddbService.GetDeployment(given.readArg)
			require.NoError(t, err)
			require.Empty(t, depl)
		})
	}

	testsError := make(map[string]given)

	testsError["when the composite is not given"] = given{
		readArg: read3,
	}

	for when, given := range testsError {
		t.Run(when, func(t *testing.T) {
			_, err := ddbService.GetDeployment(given.readArg)
			require.Error(t, err)
		})
	}

	cleanup()
}

func createDeploymentNoErrors(t *testing.T, arg CreateDeployment) {
	type given struct {
		arg CreateDeployment
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when a deployment is properly specified"] = given{
		arg: arg,
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.NoError(t, err)
		})
	}
}

func TestListDeployments(t *testing.T) {
	type given struct {
		count int
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when there are 10 items in the db"] = given{
		count: 10,
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			for i := 0; i < given.count; i++ {
				createRandomDeployment(t)
			}
			depls, err := ddbService.ListDeployments()
			require.NoError(t, err)
			require.Len(t, depls, given.count)
			for _, depl := range depls {
				require.NotEmpty(t, depl)
			}
			cleanup()
		})
	}
}

func createRandomDeployment(t *testing.T) {
	type given struct {
		arg CreateDeployment
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when a deployment is properly specified"] = given{
		arg: CreateDeployment{
			Name: util.RandomString(6),
			Version: fmt.Sprintf("%v.%v.%v", util.RandomInt(0, 99), util.RandomInt(0, 99), util.RandomInt(0, 99)),
		},
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.NoError(t, err)
		})
	}
}

func cleanup() {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Deployments"),
	}

	p := dynamodb.NewScanPaginator(ddbService.ddb, input)

	for p.HasMorePages() {
		scanOutput, err := p.NextPage(context.Background())
		if err != nil {
			log.Fatal("failed to scan:", err)
		}

		for _, item := range scanOutput.Items {
			input := &dynamodb.DeleteItemInput{
				TableName: aws.String("Deployments"),
				Key: map[string]types.AttributeValue{
					"name": item["name"],
					"version": item["version"],
				},
			}
			_, err = ddbService.ddb.DeleteItem(context.Background(), input)
			if err != nil {
				log.Fatal("failed to cleanup the db:", err)
			}
		}
	}
}
