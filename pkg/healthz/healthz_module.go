package healthz

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

// Module - healthz module
func Module(gin *gin.Engine, config *config.ConfigService, logger *logrus.Entry) {
	log := logger.WithFields(logrus.Fields{"module": "healthz"})

	NewHealthzController(gin, config)
	log.Info("healthz module loaded")
}
