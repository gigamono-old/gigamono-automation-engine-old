package engine

import (
	"github.com/sageflow/sageflow/pkg/database/models"
)

// WorkflowInstance represents a running, paused or stopped workflow instance.
type WorkflowInstance struct {
	Model    *models.WorkflowInstance
	Workflow *Workflow
}
