package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-workflow-engine/pkg/runnablesupervisor"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Resource")
	if err != nil {
		logs.FmtPrintln("Unable to initialize engine:", err)
		return
	}

	// Start main server.
	server, err := runnablesupervisor.NewRunnableSupervisor(app)
	if err != nil {
		logs.FmtPrintln("Unable to create engine server:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("Unable to listen on port specified:", err)
	}
}
