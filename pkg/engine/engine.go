package engine

import (
	"github.com/gofrs/uuid"

	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Engine represents an engine instance.
type Engine struct {
	Model *models.Engine
	DB    *database.DB
}

// NewEngine creates a new workflow engine.
func NewEngine(db *database.DB) Engine {
	// Create engine in the database.
	engine := db.CreateEngine()
	return Engine{
		Model: engine,
		DB: db,
	}
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *Engine) ExecuteWorkflow(workflowID uuid.UUID) error {
	// If it exists.
	// Execute workflow
	return nil
}
