package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-workflow-engine/pkg/apiservice"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp(inits.WorkflowEngineAPIService)
	if err != nil {
		logs.FmtPrintln("initialising api service:", err)
		return
	}

	// Start main server.
	server, err := apiservice.NewAPIService(app)
	if err != nil {
		logs.FmtPrintln("creating api service:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
