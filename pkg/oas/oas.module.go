package oas

import (
	"github.com/gin-gonic/gin"
	oas_spec "github.com/swaggo/files"
	oas "github.com/swaggo/gin-swagger"
	docs "github.com/techdecaf/k2s/v2/docs"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

// Module - oas module
func Module(gin *gin.Engine, config *config.ConfigService) (err error) {
	docs.SwaggerInfo.BasePath = config.BASE_PATH
	gin.GET("/oas/*any", oas.WrapHandler(oas_spec.Handler))
	return err
}
