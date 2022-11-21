package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/kube"
)

func Module(app *gin.Engine, k8s *kube.Service, logger *logrus.Entry) {
	log := logger.WithFields(logrus.Fields{"module": "deployments"})
	// table, err := state.NewDeploymentsTable(stream)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	deploymentService := NewDeploymentService(k8s)
	// if err := deploymentService.OnModuleInit(); err != nil {
	// 	log.Fatal(err)
	// }

	NewDeploymentController(app, deploymentService, log)

	log.Info("deployments module loaded")
}
