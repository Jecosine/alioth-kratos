//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/resolver"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	//todos []*model.Todo
}

func (r Resolver) Mutation() resolver.MutationResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Query() resolver.QueryResolver {
	//TODO implement me
	panic("implement me")
}
