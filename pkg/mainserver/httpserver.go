package mainserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-automation-engine/internal/mainserver/graphql"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gigamono/gigamono/pkg/services/rest/routes"
)

func (server *MainServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.AutomationEngine.MainServer.Ports.Public),
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
	routes.SetLocalStaticRoutes(server.GinEngine, &server.App)

	// GraphQL handler.
	graphqlHandler := graphql.Handler(&server.App)
	graphqlRoute := server.GinEngine.Group("/graphql", middleware.AuthenticateCreateUser(&server.App))
	{
		graphqlRoute.POST("/", graphqlHandler) // Handles all graphql requests.
		graphqlRoute.GET("/", graphqlHandler)  // Handles query-only graphql requests.
	}
}
