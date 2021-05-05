package runnablesupervisor

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
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
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", supervisor.Config.Services.Types.WorkflowEngine.Ports.RunnableSupervisor))
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run supervisors concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return supervisor.grpcServe(grpcListener) })
	grp.Go(func() error { return supervisor.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}
