package engine

import (
	"github.com/gigamono/gigamono/pkg/configs"
)

// Workflow represents a runnable workflow.
type Workflow struct {
	WorkflowConfig *configs.WorkflowConfig
}

// Execute starts the execution of workflow run.
func (workflow *Workflow) Execute() error {
	// TODO
	return nil
}
