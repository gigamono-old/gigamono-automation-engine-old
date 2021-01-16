package main

import (
	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sageutils/pkg/envs"
	"github.com/sageflow/sageutils/pkg/logs"

	"github.com/sageflow/sageengine/pkg/engine"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile()
	envs.LoadEnvFile()

	// Connect to database.
	db := database.Connect()

	// Start a workflow engine gRPC server.
	eng := engine.NewEngine(db)
	eng.Listen("3001") // TODO: Get from .env with a default value.
}
