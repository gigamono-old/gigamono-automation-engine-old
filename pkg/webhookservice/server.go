package webhookservice

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// WebhookService is a grpc server with an engine.
type WebhookService struct {
	inits.App
	GinEngine *gin.Engine
}

// NewWebhookService creates a new server instance.
func NewWebhookService(app inits.App) (WebhookService, error) {
	return WebhookService{
		App:       app,
		GinEngine: gin.Default(),
	}, nil
}

// Listen makes the server listen on specified port.
func (service *WebhookService) Listen() error {
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", service.Config.Services.Types.WorkflowEngine.Ports.WebhookService))
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return service.grpcServe(grpcListener) })
	grp.Go(func() error { return service.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}
