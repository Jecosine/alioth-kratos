package resolver

import (
	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

//// NewSchema creates a graphql executable schema.
//func NewSchema(client *ent.Client) graphql.ExecutableSchema {
//	return NewExecutableSchema(Config{
//		Resolvers: &Resolver{client},
//	})
//}

// NewResolver create a resolver with client
func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		client: client,
	}
}
