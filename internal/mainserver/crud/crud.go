package crud

import (
	gqlModel "github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
	"github.com/gigamono/gigamono/pkg/configs"
	controller "github.com/gigamono/gigamono/pkg/database/controllers/resource"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/services/auth"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(app *inits.App, tokens gqlModel.TokensInput, workflow *gqlModel.WorkflowInput) (string, error) {
	// TODO: Sec: Validation, Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return "", logs.NewError("unable to authenticate user", err)
	}

	// Validate workflow config.
	if _, err := configs.NewWorkflowConfig(workflow.Specification, configs.JSON); err != nil {
		return "", logs.Error("unable to load workflow config", err)
	}

	// TODO: Compile workflow config.

	id, err := controller.CreateWorkflow(&app.DB, &sessionUserID, workflow.Name, workflow.Specification)
	if err != nil {
		return "", logs.NewError("unable to create workflow", err)
	}

	return id.String(), nil
}

// ActivateWorkflow starts running a workflow.
func ActivateWorkflow(app *inits.App, tokens gqlModel.TokensInput, workflowID string) (string, error) {
	// TODO: Sec: Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return "", logs.NewError("unable to authenticate user", err)
	}

	workflowUUID, err := uuid.FromString(workflowID)
	if err != nil {
		return "", logs.NewError("unable to parse workflow id", err)
	}

	id, err := controller.ActivateWorkflow(&app.DB, &sessionUserID, &workflowUUID)
	if err != nil {
		return "", logs.NewError("unable to activate workflow", err)
	}

	// TODO: Register event -> WebhookService, PollerService or SechedulerService

	return id.String(), nil
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(app *inits.App, tokens gqlModel.TokensInput, workflowID string) (*gqlModel.Workflow, error) {
	// TODO: Sec: Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return &gqlModel.Workflow{}, logs.NewError("unable to authenticate user", err)
	}

	workflowUUID, err := uuid.FromString(workflowID)
	if err != nil {
		return &gqlModel.Workflow{}, logs.NewError("unable to parse workflow id", err)
	}

	workflow, err := controller.GetWorkflow(&app.DB, &sessionUserID, &workflowUUID)
	if err != nil {
		return &gqlModel.Workflow{}, logs.NewError("unable to get workflow", err)
	}

	return &gqlModel.Workflow{
		ID:            workflow.ID.String(),
		Name:          workflow.Name,
		Specification: workflow.Specification,
		IsActive:      &workflow.IsActive,
		CreatorID:     workflow.CreatorID.String(),
	}, nil
}
