package mainserver

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// MainServer represents the main server.
type MainServer struct {
	inits.App
	GinEngine *gin.Engine
}

// NewMainServer creates a new server.
func NewMainServer(app inits.App) (MainServer, error) {
	return MainServer{
		App:       app,
		GinEngine: gin.Default(),
	}, nil
}

// Listen makes the server listen on specified port.
func (server *MainServer) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe() })
	grp.Go(func() error { return server.httpServe() })
	return grp.Wait()
}
