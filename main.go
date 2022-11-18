package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/db"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
	"github.com/techdecaf/k2s/v2/pkg/global"
	"github.com/techdecaf/k2s/v2/pkg/healthz"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/logger"
	"github.com/techdecaf/k2s/v2/pkg/registries"
	"github.com/techdecaf/k2s/v2/pkg/traefik"
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

	configService, err := config.NewConfigService(os.Environ()...).Validate()
	if err != nil {
		panic(err)
	}
	log := logger.NewLogger(configService)

	kubeService, err := kube.NewKubeService()
	if err != nil {
		log.Fatal(err)
	}

	ddbClient, err := db.NewDDB(configService)
	if err != nil {
		log.Fatal(err)
	}

	// create new gin application
	gin.SetMode(gin.ReleaseMode)

	dependencies := global.NewDependencies(log, gin.New(), kubeService, configService, ddbClient)

	// star the application
	for item := range dependencies.OnModuleInit().Observe() {
		err := item.E
		if err != nil {
			log.Fatal(err)
		}

		services := item.V.(*global.Server)

		if _, err = services.Kube.ApplyNamespace(&coreV1.Namespace{
			ObjectMeta: metaV1.ObjectMeta{Name: services.Config.SERVICE_NAME},
		}); err != nil && !apierrors.IsAlreadyExists(err) {
			log.Fatal(err)
		}

		// middlewares
		services.Gin.Use(logger.Middleware(configService))
		// services.Gin.Use(ddtrace.Middleware(config.SERVICE_NAME))

		// modules
		healthz.Module(services.Gin, services.Config, services.Log)
		traefik.Module(services.Gin, configService, kubeService, log)
		registries.Module(services.Gin, configService, kubeService, log)
		deployments.Module(services.Gin, kubeService, log)
	}

	// CleanUp
	log.Fatal(dependencies.Gin.Run(fmt.Sprintf("0.0.0.0:%s", configService.PORT)))
}
