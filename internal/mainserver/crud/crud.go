package crud

import (
	"context"
	"errors"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(ctx context.Context, workflow *model.WorkflowInput) (string, error) {
	// TODO: Sec: Validation, Auth, Permission.
	return "", errors.New("Testing")
}

// ActivateWorkflow starts running a workflow.
func ActivateWorkflow(ctx context.Context, id string) (string, error) {
	// TODO: Sec: Auth, Permission.
	return "", errors.New("Testing")
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(ctx context.Context, id string) (*model.Workflow, error) {
	// TODO: Sec: Auth, Permission.
	return &model.Workflow{}, errors.New("Testing")
}
