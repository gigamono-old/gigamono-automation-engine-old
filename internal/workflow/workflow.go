package workflow

import (
	"github.com/google/uuid"
)

// Runnable interface represents each runnable task in a workflow.
type Runnable interface {
	run() error
}

// Workflow represents an executable pipeline of tasks.
type Workflow struct {
	ID    uuid.UUID
	Tasks []Runnable
}
