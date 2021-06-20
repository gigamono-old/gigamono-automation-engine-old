package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-automation-engine/pkg/mainserver"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp(inits.AutomationEngineMainServer)
	if err != nil {
		logs.FmtPrintln("initialising main server:", err)
		return
	}

	// Start main server.
	server, err := mainserver.NewMainServer(app)
	if err != nil {
		logs.FmtPrintln("creating main server:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
