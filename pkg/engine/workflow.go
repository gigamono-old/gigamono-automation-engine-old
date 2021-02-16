package engine

import (
	"github.com/sageflow/sageflow/pkg/configs"
	"github.com/sageflow/sageflow/pkg/database/models"
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
