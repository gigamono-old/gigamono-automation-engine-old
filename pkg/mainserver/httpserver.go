package mainserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql"
)

func (server *MainServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.Types.WorkflowEngine.PublicPorts.MainServer),
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
	graphqlHandler := graphql.Handler(&server.App)
	playgroundHandler := graphql.PlaygroundHandler()

	server.GinEngine.POST("/graphql", graphqlHandler)      // Handles all graphql requests.
	server.GinEngine.GET("/graphql", graphqlHandler)       // Handles query-only graphql requests.
	server.GinEngine.GET("/playground", playgroundHandler) // Shows playground UI.
}
