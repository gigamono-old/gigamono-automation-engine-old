package runnablesupervisor

import (
	"net"
	"net/http"
)

func (supervisor *RunnableSupervisor) httpServe(listener net.Listener) error {
	supervisor.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: supervisor.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (supervisor *RunnableSupervisor) setRoutes() {}
