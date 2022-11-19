package deployments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/global"
	"github.com/techdecaf/k2s/v2/pkg/state"
)

// NewDeploymentController function description
func NewDeploymentController(app *gin.Engine, deploymentSvc *DeploymentService) *DeploymentController {
	controller := &DeploymentController{deploy: deploymentSvc}

	// register routes
	router := app.Group("/deployments")

	router.POST("", controller.CreateDeployment)
	router.GET("/:name/:version", controller.GetDeployment)
	router.GET("", controller.ListDeployments)
	// router.DELETE("/:id", controller.DeleteDeployment)

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
	resp, err := t.deploy.ListDeployments()
	if err != nil {
		t.deploy.log.Error(err)
		global.GinerateError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, resp)
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

// @Summary list deployed services
// @Description list deployed services
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments [GET]
// ListDeployments method
func (t *DeploymentController) CreateDeployment(context *gin.Context) {
	var deploymentReq createDeploymentRequest

	if err := context.ShouldBind(&deploymentReq); err != nil {
		global.GinerateError(context, global.BadRequestError(err))
		return
	}

	deployment := state.DeploymentDTO{
		Name:        deploymentReq.Name,
		Image:       deploymentReq.Image,
		Version:     deploymentReq.Version,
		Environment: deploymentReq.Environment,
	}

	err := t.deploy.CreateDeployment(&deployment)
	if err != nil {
		t.deploy.log.Error(err)
		global.GinerateError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, nil)
}

func (t *DeploymentController) GetDeployment(context *gin.Context) {
	var deploymentReq readDeploymentRequest
	if err := context.ShouldBindUri(&deploymentReq); err != nil {
		global.GinerateError(context, global.BadRequestError(err))
		return
	}

	depl := state.DeploymentDTO{
		Name:    deploymentReq.Name,
		Version: deploymentReq.Version,
	}

	resp, err := t.deploy.GetDeployment(&depl)
	if err != nil {
		t.deploy.log.Error(err)
		global.GinerateError(context, global.InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, resp)
}
