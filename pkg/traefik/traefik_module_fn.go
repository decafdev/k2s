package traefik

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/kube"
)

// Module - traefik module
func Module(app *gin.Engine, config *config.ConfigService, kube *kube.Service, logger *logrus.Entry) {
	log := logger.WithField("module", "traefik")

	traefikService := &TraefikService{config: config, k8s: kube, log: log}
	if err := traefikService.OnModuleInit(); err != nil {
		log.Fatal(err)
	}

	NewTraefikController(app, traefikService)
	log.Info("traefik module loaded")
}
