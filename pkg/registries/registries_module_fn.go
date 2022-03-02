package registries

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/kube"
)

// Module registries module
func Module(app *gin.Engine, config *config.ConfigService, k8s *kube.Service, logger *logrus.Entry) {
	log := logger.WithFields(logrus.Fields{"module": "registries"})

	registryService := &RegistryService{config: config, k8s: k8s, log: log}

	if err := registryService.OnModuleInit(); err != nil {
		log.Fatal(err)
	}

	NewRegistryController(app, registryService)

	log.Info("registries module loaded")
}
