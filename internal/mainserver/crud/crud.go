package crud

import (
	"fmt"

	gqlModel "github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
	"github.com/gigamono/gigamono/pkg/configs"
	controller "github.com/gigamono/gigamono/pkg/database/controllers/resource"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/auth"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(app *inits.App, tokens gqlModel.TokensInput, workflow *gqlModel.WorkflowInput) (string, error) {
	// TODO: Sec: Validation, Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return "", fmt.Errorf("unable to authenticate user: %v", err)
	}

	specification, err := configs.NewWorkflowConfig(workflow.Specification, configs.JSON)
	if err != nil {
		return "", fmt.Errorf("unable to load workflow config: %v", err)
	}

	return controller.CreateWorkflow(&app.DB, &sessionUserID, workflow.Name, specification)
}

// ActivateWorkflow starts running a workflow.
func ActivateWorkflow(app *inits.App, tokens gqlModel.TokensInput, id string) (string, error) {
	// TODO: Sec: Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return "", fmt.Errorf("unable to authenticate user: %v", err)
	}

	workflowID, err := uuid.FromString(id)
	if err != nil {
		return "", fmt.Errorf("unable to activate workflow config: %v", err)
	}

	return controller.ActivateWorkflow(&app.DB, &sessionUserID, &workflowID)
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(app *inits.App, tokens gqlModel.TokensInput, id string) (*gqlModel.Workflow, error) {
	// TODO: Sec: Auth, Permission.
	sessionUserID, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		return &gqlModel.Workflow{}, fmt.Errorf("unable to authenticate user: %v", err)
	}

	workflowID, err := uuid.FromString(id)
	if err != nil {
		return &gqlModel.Workflow{}, fmt.Errorf("unable to activate workflow config: %v", err)
	}

	workflow, err := controller.GetWorkflow(&app.DB, &sessionUserID, &workflowID)
	if err != nil {
		return &gqlModel.Workflow{}, fmt.Errorf("unable to get workflow: %v", err)
	}

	specification, err := workflow.Specification.JSON()
	if err != nil {
		return &gqlModel.Workflow{}, fmt.Errorf("unable convert workflow to json: %v", err)
	}

	return &gqlModel.Workflow{
		ID:            workflow.ID.String(),
		Name:          workflow.Name,
		Specification: specification,
		IsActive:      &workflow.IsActive,
		CreatorID:     workflow.CreatorID.String(),
	}, nil
}
