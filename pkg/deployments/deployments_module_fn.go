package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/state"
	"github.com/techdecaf/k2s/v2/pkg/streams"
)

func Module(app *gin.Engine, stream *streams.Client, k8s *kube.Service, logger *logrus.Entry) {
	log := logger.WithFields(logrus.Fields{"module": "deployments"})
	table, err := state.NewDeploymentsTable(stream)

	if err != nil {
		log.Fatal(err)
	}

	deploymentService := NewDeploymentService(table, k8s, log)
	if err := deploymentService.OnModuleInit(); err != nil {
		log.Fatal(err)
	}

	NewDeploymentController(app, deploymentService)

	log.Info("deployments module loaded")
}
