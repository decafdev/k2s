package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDeployment(t *testing.T) {
	type given struct {
		arg CreateDeployment
	}

	testsNoErrors := make(map[string]given)

	testsNoErrors["when a deployment is properly specified"] = given{
		arg: CreateDeployment{
			Image: "repo/image",
			Version: "0.0.0",
		},
	}

	testsNoErrors["when the user requests to deploy again"] = given{
		arg: CreateDeployment{
			Image: "repo/image",
			Version: "0.0.0",
		},
	}

	for when, given := range testsNoErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.NoError(t, err)
		})
	}

	testsErrors := make(map[string]given)

	testsErrors["when no deployment is specified"] = given{
		arg: CreateDeployment{},
	}

	testsErrors["when a deployment is improperly specified"] = given{
		arg: CreateDeployment{
			Version: "monotonic",
		},
	}

	for when, given := range testsErrors {
		t.Run(when, func(t *testing.T) {
			err := ddbService.CreateDeployment(given.arg)
			require.Error(t, err)
		})
	}
}