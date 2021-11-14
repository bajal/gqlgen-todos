package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/bajal/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// To track our state. This is an in-memory todos list
type Resolver struct {
	todos []*model.Todo
}
