package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/config"
)

// NewHealthzController function description
func NewHealthzController(app *gin.Engine, config *config.ConfigService) *HealthzController {
	controller := &HealthzController{
		config: config,
	}

	// register routes
	router := app.Group("/healthz")

	router.GET("", controller.GetHealthz)

	return controller
}

// HealthzController struct
type HealthzController struct {
	config *config.ConfigService
}

// @Summary healthz
// @Description healthz
// @Accept application/json
// @Produce json
// @Success 200 {object} HealthDTO true
// @Router /healthz [GET]
// GetHealthz method
func (t *HealthzController) GetHealthz(context *gin.Context) {
	context.JSON(http.StatusOK, &HealthDTO{
		Name:     t.config.SERVICE_NAME,
		Version:  t.config.VERSION,
		Hostname: context.Request.Host,
	})
}
