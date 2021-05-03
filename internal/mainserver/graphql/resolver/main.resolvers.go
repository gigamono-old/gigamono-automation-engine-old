package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/crud"
	generated1 "github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/generated"
	model1 "github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
)

func (r *mutationResolver) CreateWorkflow(ctx context.Context, workflow model1.WorkflowInput) (string, error) {
	return crud.CreateWorkflow(ctx, &workflow)
}

func (r *mutationResolver) ActivateWorkflow(ctx context.Context, id string) (string, error) {
	return crud.ActivateWorkflow(ctx, id)
}

func (r *queryResolver) GetWorkflow(ctx context.Context, id string) (*model1.Workflow, error) {
	return crud.GetWorkflow(ctx, id)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
