package webhookservice

import (
	"fmt"
	"net"
	"net/http"
)

func (service *WebhookService) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(
			":",
			service.Config.Services.Types.WorkflowEngine.PublicPorts.WebhookService,
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

func (service *WebhookService) setRoutes() {}
