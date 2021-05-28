package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/crud"
	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/generated"
	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
)

func (r *mutationResolver) CreateWorkflow(ctx context.Context, workflow model.WorkflowInput) (string, error) {
	return crud.CreateWorkflow(ctx, r.App, &workflow)
}

func (r *mutationResolver) ActivateWorkflow(ctx context.Context, workflowID string) (string, error) {
	return crud.ActivateWorkflow(ctx, r.App, workflowID)
}

func (r *queryResolver) GetWorkflow(ctx context.Context, workflowID string) (*model.Workflow, error) {
	return crud.GetWorkflow(ctx, r.App, workflowID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
