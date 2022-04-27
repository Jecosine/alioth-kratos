package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/model"
)

func (r *mutationResolver) CreateAnnouncement(ctx context.Context, announcementInput model.AnnouncementInput) (*ent.Announcement, error) {
	announcement, err := r.client.Announcement.Create().
		SetContent(announcementInput.Content).
		SetTitle(announcementInput.Title).
		SetCreatedTime(time.Now()).
		SetAuthorID(announcementInput.UserID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return announcement, nil
}

func (r *mutationResolver) UpdateAnnouncement(ctx context.Context, announcementInput model.AnnouncementInputWithID) (*ent.Announcement, error) {
	save, err := r.client.Announcement.
		UpdateOneID(announcementInput.UserID).
		SetContent(announcementInput.Content).
		SetTitle(announcementInput.Title).
		SetModifiedTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) DeleteAnnouncement(ctx context.Context, announcementID int64) (*model.Reply, error) {
	err := r.client.Announcement.DeleteOneID(announcementID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Deleted",
		Data:    nil,
	}, nil
}

func (r *queryResolver) GetAnnouncements(ctx context.Context, userID int64) ([]*ent.Announcement, error) {
	// get all announcement concerned with user: all team, global message
	// get all team
	teams, err := r.client.User.GetX(ctx, userID).Teams(ctx)
	if err != nil {
		return nil, err
	}
	// TODO: pagination and select recent announcements
	var allAnnouncement []*ent.Announcement
	for _, team := range teams {
		all, err := r.client.Team.QueryAnnouncements(team).All(ctx)
		if err != nil {
			return nil, err
		}
		allAnnouncement = append(allAnnouncement, all...)
	}
	return allAnnouncement, nil
}
