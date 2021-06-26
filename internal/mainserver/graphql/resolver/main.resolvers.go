package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gigamono/gigamono-automation-engine/internal/mainserver/crud"
	"github.com/gigamono/gigamono-automation-engine/internal/mainserver/graphql/generated"
	"github.com/gigamono/gigamono-automation-engine/internal/mainserver/graphql/model"
)

func (r *mutationResolver) CreateWorkflow(ctx context.Context, automationID string, workflow model.WorkflowInput) (*model.Workflow, error) {
	return crud.CreateWorkflow(ctx, r.App, &automationID, &workflow)
}

func (r *mutationResolver) PatchWorkflowSpecification(ctx context.Context, id string, patch string) (*model.Workflow, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Workflow(ctx context.Context, id string) (*model.Workflow, error) {
	return crud.GetWorkflow(ctx, r.App, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
