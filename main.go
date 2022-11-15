package main

import (
	"fmt"
	"os"

	"github.com/techdecaf/k2s/v2/pkg/api"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/traefik"
	"github.com/techdecaf/k2s/v2/pkg/util"
	coreV1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VERSION - This is converted to the git tag at compile time using tasks run build command
var VERSION = "0.0.0"
var SERVICE_NAME = "k2s-operator"

// @title k2s operator
// @version 2.0
// @description staggeringly simple and opinionated kubernetes deployments
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// set version environment variable
	os.Setenv("SERVICE_NAME", SERVICE_NAME)
	os.Setenv("VERSION", VERSION)

	config, err := util.NewConfig(os.Environ()...).Validate()
	if err != nil {
		panic(err)
	}

	logger := util.NewLogger(config)

	kubeClient, err := kube.NewKubeClient()
	if err != nil {
		logger.Fatal(err)
	}

	_, err = kubeClient.ApplyNamespace(&coreV1.Namespace{ObjectMeta: metaV1.ObjectMeta{Name: config.SERVICE_NAME}})
	if err != nil && !apierrors.IsAlreadyExists(err) {
		logger.Fatal(err)
	}

	err = traefik.StartTraefik(kubeClient, config, logger)
	if err != nil {
		logger.Fatal(err)
	}

	deploymentService := deployments.NewDeploymentService(kubeClient, logger)

	server := api.NewServer(config, logger, deploymentService)

	address := fmt.Sprintf("0.0.0.0:%s", config.PORT)

	err = server.Start(address)
	if err != nil {
		logger.Fatal(err)
	}
}
