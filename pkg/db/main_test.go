package db

import (
	"log"
	"os"
	"testing"

	"github.com/techdecaf/k2s/v2/pkg/config"
)

var ddbService *DDBService

func TestMain(m *testing.M) {
	config, err := config.NewConfigService(os.Environ()...).Validate()
	if err != nil {
		panic(err)
	}

	ddbService, err = NewDDB(config)
	if err != nil {
		log.Fatal("did not setup ddb:", err)
	}

	os.Exit(m.Run())
}