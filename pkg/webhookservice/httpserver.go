package webhookservice

import (
	"net"
	"net/http"
)

func (service *WebhookService) httpServe(listener net.Listener) error {
	service.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: service.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (service *WebhookService) setRoutes() {}
