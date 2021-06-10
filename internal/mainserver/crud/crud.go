package crud

import (
	"context"

	gqlModel "github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
	"github.com/gigamono/gigamono/pkg/configs"
	controller "github.com/gigamono/gigamono/pkg/database/controllers/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(ctx context.Context, app *inits.App, workflow *gqlModel.WorkflowInput) (string, error) {
	// TODO: Sec: Validation, Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).UserID

	// Validate workflow config.
	if _, err := configs.NewWorkflowConfig(workflow.Specification, configs.JSON); err != nil {
		panic(errs.NewSystemError(
			messages.Error["workflow-config"].(string),
			"loading workflow config",
			err,
		))
	}

	// TODO: Compile workflow config.

	// TODO: Save workflow assets.

	// Create the workflow in db.
	id, err := controller.CreateWorkflow(&app.DB, &userID, workflow.Name)
	if err != nil {
		panic(errs.NewSystemError("", "creating workflow", err))
	}

	return id.String(), nil
}

// ActivateWorkflow starts running a workflow.
func ActivateWorkflow(ctx context.Context, app *inits.App, workflowID string) (string, error) {
	// TODO: Sec: Validation, Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).UserID

	workflowUUID, err := uuid.FromString(workflowID)
	if err != nil {
		panic(err)
	}

	// Activate the workflow in db.
	id, err := controller.ActivateWorkflow(&app.DB, &userID, &workflowUUID)
	if err != nil {
		panic(errs.NewSystemError("", "activating workflow", err))
	}

	// TODO: Register event -> WebhookService, PollerService or SechedulerService

	return id.String(), nil
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(ctx context.Context, app *inits.App, workflowID string) (*gqlModel.Workflow, error) {
	// TODO: Sec: Validation, Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).UserID

	workflowUUID, err := uuid.FromString(workflowID)
	if err != nil {
		panic(err)
	}

	// Get the workflow from db.
	workflow, err := controller.GetWorkflow(&app.DB, &userID, &workflowUUID)
	if err != nil {
		panic(errs.NewSystemError("", "getting workflow", err))
	}

	return &gqlModel.Workflow{
		ID:   workflow.ID.String(),
		Name: workflow.Name,
		// Specification: workflow.Specification,
		IsActive:  &workflow.IsActive,
		CreatorID: workflow.CreatorID.String(),
	}, nil
}
