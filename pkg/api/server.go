package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/deployments"
	"github.com/techdecaf/k2s/v2/pkg/util"
)

// Server struct
type Server struct {
	routerService *gin.Engine
	config *util.Config
	logger *logrus.Entry
	deploymentService *deployments.DeploymentService
}

// // OnModuleInit method
// func (t *Server) OnModuleInit() rxgo.Observable {
// 	return rxgo.Just(t)()
// }

func NewServer(config *util.Config, log *logrus.Entry, depl *deployments.DeploymentService) *Server {
	server := &Server{
		config: config,
		logger: log,
		deploymentService: depl,
	}

	server.setupRouter()
	
	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()

	healthRouter := router.Group("/healthz")

	healthRouter.GET("", s.getHealthz)

	deploymentsRouter := router.Group("/deployments")

	deploymentsRouter.POST("", s.createDeployment)

	s.routerService = router
}

func (s *Server) Start(address string) error {
	return s.routerService.Run(address)
}
