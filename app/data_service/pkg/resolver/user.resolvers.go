package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Jecosine/alioth-kratos/app/data_service/pkg/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Ping(ctx context.Context, payload model.PingInput) (*string, error) {
	return &payload.Msg, nil
	//panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PingAuth(ctx context.Context, payload model.PingInput) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
