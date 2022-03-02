package deployments

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewDeploymentController function description
func NewDeploymentController(app *gin.Engine, deploymentSvc *DeploymentService) *DeploymentController {
	controller := &DeploymentController{deploy: deploymentSvc}

	// register routes
	router := app.Group("/deployments")

	router.GET("", controller.ListDeployments)

	return controller
}

// DeploymentController struct
type DeploymentController struct {
	deploy *DeploymentService
}

// @Summary list deployed services
// @Description list deployed services
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments [GET]
// ListDeployments method
func (t *DeploymentController) ListDeployments(context *gin.Context) {
	context.JSON(http.StatusOK, []DeploymentDTO{{}})
}
