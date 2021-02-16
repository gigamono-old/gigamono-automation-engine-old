package engine

import (
	"github.com/gofrs/uuid"

	"github.com/sageflow/sageflow/pkg/database/models"
	"github.com/sageflow/sageflow/pkg/inits"
)

// Engine represents an engine instance.
type Engine struct {
	*inits.App
	Model models.Engine
}

// NewEngine creates a new workflow engine.
func NewEngine(app *inits.App) (Engine, error) {
	// Create engine in the database.
	engine, err := app.DB.CreateEngine()
	if err != nil {
		return Engine{}, err
	}

	return Engine{
		App:   app,
		Model: engine,
	}, nil
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *Engine) ExecuteWorkflow(workflowID uuid.UUID) error {
	// If it exists.

	// Execute workflow
	return nil
}
