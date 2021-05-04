package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/crud"
	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/generated"
	"github.com/gigamono/gigamono-workflow-engine/internal/mainserver/graphql/model"
)

func (r *mutationResolver) CreateWorkflow(ctx context.Context, tokens model.TokensInput, workflow model.WorkflowInput) (string, error) {
	return crud.CreateWorkflow(r.App, tokens, &workflow)
}

func (r *mutationResolver) ActivateWorkflow(ctx context.Context, tokens model.TokensInput, id string) (string, error) {
	return crud.ActivateWorkflow(r.App, tokens, id)
}

func (r *queryResolver) GetWorkflow(ctx context.Context, tokens model.TokensInput, id string) (*model.Workflow, error) {
	return crud.GetWorkflow(r.App, tokens, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
