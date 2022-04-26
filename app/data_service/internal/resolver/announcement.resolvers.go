package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
)

func (r *announcementResolver) Author(ctx context.Context, obj *ent.Announcement) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Announcement returns AnnouncementResolver implementation.
func (r *Resolver) Announcement() AnnouncementResolver { return &announcementResolver{r} }

type announcementResolver struct{ *Resolver }
