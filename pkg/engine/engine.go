package engine

import (
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
