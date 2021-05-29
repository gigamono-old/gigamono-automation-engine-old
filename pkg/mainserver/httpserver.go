package mainserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql"
	"github.com/gigamono/gigamono/pkg/services/rest"
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
	// Set local static routes if specified.
	rest.SetLocalStaticRoutes(server.GinEngine, &server.App)

	// Handlers.
	graphqlHandler := graphql.Handler(&server.App)
	server.GinEngine.POST("/graphql", graphqlHandler) // Handles all graphql requests.
	server.GinEngine.GET("/graphql", graphqlHandler)  // Handles query-only graphql requests.
}
