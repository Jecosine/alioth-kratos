package service

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/resolver"
)

func NewResolver(drv *sql.Driver) *resolver.Resolver {
	client := ent.NewClient(ent.Driver(drv))
	err := client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return resolver.NewResolver(client)
}

func NewEntClient(drv *sql.Driver) *ent.Client {
	client := ent.NewClient(ent.Driver(drv))
	err := client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
