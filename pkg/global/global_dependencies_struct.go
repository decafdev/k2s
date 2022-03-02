package global

import (
	"github.com/gin-gonic/gin"
	"github.com/reactivex/rxgo/v2"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/kube"
	"github.com/techdecaf/k2s/v2/pkg/streams"
)

// Dependencies struct
type Dependencies struct {
	Log     *logrus.Entry
	Gin     *gin.Engine
	Kube    *kube.Service
	Config  *config.ConfigService
	Streams *streams.Client
}

// OnModuleInit method
func (t *Dependencies) OnModuleInit() rxgo.Observable {
	return rxgo.Just(t)()
}

// NewDependencies function description
func NewDependencies(
	Log *logrus.Entry,
	Gin *gin.Engine,
	Kube *kube.Service,
	Config *config.ConfigService,
	Streams *streams.Client,
) *Dependencies {

	return &Dependencies{
		Log:     Log,
		Gin:     Gin,
		Kube:    Kube,
		Config:  Config,
		Streams: Streams,
	}
}
