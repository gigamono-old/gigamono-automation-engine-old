package main

import (
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/envs"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageengine/pkg/server"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile() // TODO. logs.SetStatusLogFile(config.Logging.Status.Filepath)
	envs.LoadEnvFile()      // TODO. Remove!

	// Connect to database.
	db := database.Connect() // TODO. database.Connect(config.db)

	// Start a workflow engine gRPC server.
	eng := server.NewEngineServer(db) // TODO. engine.NewEngine(db, config)
	eng.Listen("3001") // TODO. database.Connect(config.Server.Engine.Port)
}
