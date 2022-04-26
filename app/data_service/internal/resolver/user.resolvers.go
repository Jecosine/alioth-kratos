package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Ping(ctx context.Context, payload model.PingInput) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PingAuth(ctx context.Context, payload model.PingInput) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentUser(ctx context.Context, id string) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *userResolver) Managed(ctx context.Context, obj *ent.User) ([]*ent.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Announcements(ctx context.Context, obj *ent.User) ([]*ent.Announcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Records(ctx context.Context, obj *ent.User) ([]*ent.JudgeRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
//func (r *userResolver) CreatedTime(ctx context.Context, obj *ent.User) (int, error) {
//	panic(fmt.Errorf("not implemented"))
//}
