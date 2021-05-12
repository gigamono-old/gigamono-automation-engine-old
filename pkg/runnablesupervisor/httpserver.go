package runnablesupervisor

import (
	"fmt"
	"net"
	"net/http"
)

func (supervisor *RunnableSupervisor) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(
			":",
			supervisor.Config.Services.WorkflowEngine.RunnableSupervisor.Ports.Public,
		),
	)
	if err != nil {
		return err
	}

	supervisor.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: supervisor.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (supervisor *RunnableSupervisor) setRoutes() {}
