package engine

import (
	"github.com/gigamono/gigamono/pkg/configs"
	models "github.com/gigamono/gigamono/pkg/database/models/resource"
)

// Workflow represents a runnable workflow.
type Workflow struct {
	Model          *models.Workflow
	WorkflowConfig *configs.WorkflowConfig
}

// Execute starts the execution of workflow run.
func (workflow *Workflow) Execute() error {
	// TODO
	return nil
}
