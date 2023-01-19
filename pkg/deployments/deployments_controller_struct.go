package deployments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/global"
)

// DeploymentController struct
type DeploymentController struct {
	deploy *DeploymentService
}

// NewDeploymentController function description
func NewDeploymentController(app *gin.Engine, deploymentSvc *DeploymentService) *DeploymentController {
	controller := &DeploymentController{deploy: deploymentSvc}

	// register routes
	router := app.Group("/deployments")

	router.GET("", controller.ListDeployments)
	router.GET("/:name/:version", controller.GetDeployment)
	router.POST("", controller.CreateDeployment)
	// router.DELETE("/:id", controller.DeleteDeployment)

	return controller
}

// @Summary list deployed services
// @Description list deployed services
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments [GET]
// ListDeployments method
func (t *DeploymentController) CreateDeployment(context *gin.Context) {
	var deployment DeploymentDTO

	if err := context.ShouldBind(&deployment); err != nil {
		global.GinError(context, global.BadRequestError(err))
		return
	}

	err := t.deploy.CreateDeployment(&deployment)
	if err != nil {
		t.deploy.log.Error(err)
		global.GinError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, nil)
}

// @Summary list deployed services
// @Description list deployed services
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments [GET]
// ListDeployments method
func (t *DeploymentController) ListDeployments(context *gin.Context) {
	res, err := t.deploy.ListDeployments()
	if err != nil {
		t.deploy.log.Error(err)
		global.GinError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, res)
}

// @Summary get a single deployed service
// @Description get a single deployed service
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments/name/version [GET]
// ListDeployments method
func (t *DeploymentController) GetDeployment(context *gin.Context) {
	res, err := t.deploy.GetDeployment(context.Param("name"), context.Param("version"))

	if err != nil {
		t.deploy.log.Error(err)
		global.GinError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, res)
}

// @Summary delete a deployment
// @Description delete a deployment
// @Accept application/json
// @Produce json
// @Param BodyDTO body BodyDTO true "request body"
// @Success 200 {object} BodyDTO
// @Router /deployments [GET]
// func (t *DeploymentController) DeleteDeployment(context *gin.Context) {
// 	if err := t.deploy.DeleteDeployment(context.Param("id")); err != nil {
// 		global.GinerateError(context, global.InternalServerError(err))
// 	}
// 	context.JSON(http.StatusAccepted, "success")
// }
