package main

import (
	"github.com/sageflow/sagedb"
	"github.com/sageflow/sageengine"
	"github.com/sageflow/sageutils"
)

func main() {
	// Set up log status file and load .env file.
	sageutils.SetStatusLogFile()
	sageutils.LoadEnvFile()

	// Connect to database.
	db := sagedb.Connect()

	// TODO: Get from .env with a default value.
	const port = "3001"

	// Start a workflow engine gRPC server.
	engine := sageengine.NewEngine(db)
	engine.Start(port)
}
