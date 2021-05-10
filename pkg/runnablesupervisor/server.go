package runnablesupervisor

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// RunnableSupervisor is a grpc server with an engine.
type RunnableSupervisor struct {
	inits.App
	GinEngine *gin.Engine
}

// NewRunnableSupervisor creates a new supervisor instance.
func NewRunnableSupervisor(app inits.App) (RunnableSupervisor, error) {
	return RunnableSupervisor{
		App:       app,
		GinEngine: gin.Default(),
	}, nil
}

// Listen makes the supervisor listen on specified port.
func (supervisor *RunnableSupervisor) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return supervisor.grpcServe() })
	grp.Go(func() error { return supervisor.httpServe() })
	return grp.Wait()
}
