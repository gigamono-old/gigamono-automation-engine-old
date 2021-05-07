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
		logs.FmtPrintln("initialising runnable supervisor:", err)
		return
	}

	// Start main server.
	server, err := runnablesupervisor.NewRunnableSupervisor(app)
	if err != nil {
		logs.FmtPrintln("creating runnable supervisor:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
