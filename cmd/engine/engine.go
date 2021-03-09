package main

import (
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageengine/pkg/server"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Resource")
	if err != nil {
		logs.FmtPrintln("Unable to initialize engine:", err)
		return
	}

	// Start an engine gRPC server.
	server, err := server.NewEngineServer(app)
	if err != nil {
		logs.FmtPrintln("Unable to create engine server:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("Unable to listen on port specified:", err)
	}
}
