package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/boomitti/gqlgen-todos/graph/model"
)

// CreateWorkflow is the resolver for the createWorkflow field.
func (r *mutationResolver) CreateWorkflow(ctx context.Context, input model.NewWorkflow) (*model.Workflow, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	workflow := &model.Workflow{
		Name: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
	}
	r.workflows = append(r.workflows, workflow)
	return workflow, nil
}

// Workflows is the resolver for the workflows field.
func (r *queryResolver) Workflows(ctx context.Context) ([]*model.Workflow, error) {
	return r.workflows, nil
}