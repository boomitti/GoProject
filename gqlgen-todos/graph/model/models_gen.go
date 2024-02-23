// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewWorkflow struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

type Query struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Workflow struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
