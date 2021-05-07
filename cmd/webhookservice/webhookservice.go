package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-workflow-engine/pkg/webhookservice"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Resource")
	if err != nil {
		logs.FmtPrintln("initialising webhook service:", err)
		return
	}

	// Start main server.
	server, err := webhookservice.NewWebhookService(app)
	if err != nil {
		logs.FmtPrintln("creating webhook service:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
