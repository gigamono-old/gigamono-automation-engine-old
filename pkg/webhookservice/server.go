package webhookservice

import (
	"github.com/gin-gonic/gin"
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
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return service.grpcServe() })
	grp.Go(func() error { return service.httpServe() })
	return grp.Wait()
}
