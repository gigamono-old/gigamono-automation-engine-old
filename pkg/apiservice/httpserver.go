package apiservice

import (
	"fmt"
	"net"
	"net/http"
)

func (service *APIService) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(
			":",
			service.Config.Services.AutomationEngine.APIService.Ports.Public,
		),
	)
	if err != nil {
		return err
	}

	service.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: service.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (service *APIService) setRoutes() {}
