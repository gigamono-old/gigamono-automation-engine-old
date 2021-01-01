package main

import (
	"fmt"
	"github.com/sageflow/sageutils"
	"github.com/sageflow/sagedb"
)

func main() {
	sageutils.SetStatusLogFile() // Set where status log output goes.

	sageutils.LoadEnvFile() // Load env file.

	// processArgs() // Handle CLI arguments.

	db := sagedb.Connect() // Set up database.

	defer db.Close() // Close database on exit.

	fmt.Print("Engine started")
}
