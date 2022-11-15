package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
)

// @Summary list deployed services
// @Description list deployed services
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]string true
// @Router /deployments [GET]
// ListDeployments method
func (s *Server) listDeployments(context *gin.Context) {
	// context.JSON(http.StatusOK, []DeploymentDTO{{}})
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
func (s *Server) createDeployment(context *gin.Context) {
	var deployment createDeploymentRequest

	if err := context.ShouldBind(&deployment); err != nil {
		GinerateError(context, BadRequestError(err))
		return
	}

	depl := &deployments.CreateDeploymentModel{
		Name: deployment.Name,
		Image: deployment.Image,
		Version: deployment.Version,
		Environment: deployment.Environment,
	}

	err := s.deploymentService.CreateDeployment(depl)
	if err != nil {
		s.logger.Error(err)
		GinerateError(context, InternalServerError(err))
		return
	}

	context.JSON(http.StatusOK, &createDeploymentResponse{Name: deployment.Name, Image: deployment.Image, Version: deployment.Version, Environment: deployment.Environment})
}
