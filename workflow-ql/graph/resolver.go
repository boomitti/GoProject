// graph/resolver.go
package graph

import (
	"context"
	"log"
)

type Resolver struct{}

func (r *Resolver) Workflows(ctx context.Context) ([]Workflow, error) {
	var workflows []Workflow

	err := db.NewSelect().Model(&workflows).Scan(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return workflows, nil
}

func (r *Resolver) CreateWorkflow(ctx context.Context) (Workflow, error) {
	workflow := Workflow{Name: "New Workflow"}

	_, err := db.NewInsert().Model(&workflow).Exec(ctx)
	if err != nil {
		log.Println(err)
		return Workflow{}, err
	}

	return workflow, nil
}
