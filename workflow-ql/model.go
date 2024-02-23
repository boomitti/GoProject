// model.go
package main

import "time"

type Model struct {
	ID        int64     `bun:",pk,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Workflow struct {
	Model
	Name string `bun:",notnull"`
}
