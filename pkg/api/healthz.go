package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary healthz
// @Description healthz
// @Accept application/json
// @Produce json
// @Success 200 {object} HealthDTO true
// @Router /healthz [GET]
// GetHealthz method
func (s *Server) getHealthz(context *gin.Context) {
	context.JSON(http.StatusOK, &getHealthResponse{
		Name:     s.config.SERVICE_NAME,
		Version:  s.config.VERSION,
		Hostname: context.Request.Host,
	})
}
