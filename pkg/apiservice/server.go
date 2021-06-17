package apiservice

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// APIService is a grpc server with an engine.
type APIService struct {
	inits.App
	GinEngine *gin.Engine
}

// NewAPIService creates a new server instance.
func NewAPIService(app inits.App) (APIService, error) {
	return APIService{
		App:       app,
		GinEngine: gin.Default(),
	}, nil
}

// Listen makes the server listen on specified port.
func (service *APIService) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return service.grpcServe() })
	grp.Go(func() error { return service.httpServe() })
	return grp.Wait()
}
