package engine

import (
	"github.com/gofrs/uuid"

	"github.com/gigamono/gigamono/pkg/inits"
)

// WorkflowEngine represents an engine instance.
type WorkflowEngine struct {
	*inits.App
}

// NewWorkflowEngine creates a new workflow engine.
func NewWorkflowEngine(app *inits.App) (WorkflowEngine, error) {
	return WorkflowEngine{
		App: app,
	}, nil
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *WorkflowEngine) ExecuteWorkflow(workflowID uuid.UUID) error {
	// If it exists.

	// Execute workflow
	return nil
}
