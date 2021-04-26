package engine

import (
	"github.com/gofrs/uuid"

	controllers "github.com/gigamono/gigamono/pkg/database/controllers/resource"
	models "github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/inits"
)

// WorkflowEngine represents an engine instance.
type WorkflowEngine struct {
	*inits.App
	Model models.Engine
}

// NewWorkflowEngine creates a new workflow engine.
func NewWorkflowEngine(app *inits.App) (WorkflowEngine, error) {
	// Create engine in the database.
	engine, err := controllers.CreateEngine(&app.DB)
	if err != nil {
		return WorkflowEngine{}, err
	}

	return WorkflowEngine{
		App:   app,
		Model: engine,
	}, nil
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *WorkflowEngine) ExecuteWorkflow(workflowID uuid.UUID) error {
	// If it exists.

	// Execute workflow
	return nil
}
