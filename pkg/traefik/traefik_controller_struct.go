package traefik

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/global"
)

// NewTraefikController function description
func NewTraefikController(app *gin.Engine, traefik *TraefikService) *TraefikController {
	controller := &TraefikController{
		traefik: traefik,
	}

	router := app.Group("/traefik")
	router.GET("/config", controller.GetConfig)
	// router.PUT("/config", controller.GetConfig)
	return controller
}

// TraefikController struct
type TraefikController struct {
	traefik *TraefikService
}

// @Summary returns traefik config file information
// @Description returns traefik config file information
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /traefik/config [GET]
// GetConfig method
func (t *TraefikController) GetConfig(context *gin.Context) {
	config, err := t.traefik.GetTraefikConfig()
	if err != nil {
		global.GinerateError(context, global.KubeError(err))
		return
	}
	context.JSON(http.StatusOK, config["traefik-middlewares.json"])
	// context.JSON(http.StatusOK, config)
}
