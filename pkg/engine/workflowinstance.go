package engine

import (
	models "github.com/sageflow/sageflow/pkg/database/models/resource"
)

// WorkflowInstance represents a running, paused or stopped workflow instance.
type WorkflowInstance struct {
	Model    *models.WorkflowInstance
	Workflow *Workflow
}
