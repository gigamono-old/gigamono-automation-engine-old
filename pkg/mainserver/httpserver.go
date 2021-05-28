package mainserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql"
	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
)

func (server *MainServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.WorkflowEngine.MainServer.Ports.Public),
	)
	if err != nil {
		return err
	}

	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (server *MainServer) setRoutes() {
	// Depending on service config, create a local static folder for serving serverless files.
	if server.Config.Filestore.Serverless.Kind == configs.Local {
		// TODO: Permission middleware.
		// Authenticate session user.
		workflowStaticRoute := server.GinEngine.Group("/serverless", middleware.Authenticate(&server.App))
		workflowStaticRoute.StaticFS("/", http.Dir(server.Config.Filestore.Serverless.Path))
	}

	// Handlers.
	graphqlHandler := graphql.Handler(&server.App)
	server.GinEngine.POST("/graphql", graphqlHandler) // Handles all graphql requests.
	server.GinEngine.GET("/graphql", graphqlHandler)  // Handles query-only graphql requests.
}
