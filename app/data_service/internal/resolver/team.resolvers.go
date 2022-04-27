package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/team"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/model"
)

func (r *mutationResolver) AddTeam(ctx context.Context, team model.TeamInput) (*ent.Team, error) {
	save, err := r.client.Team.Create().
		SetCreatedTime(time.Now()).
		SetName(team.Name).
		SetDescription(team.Description).
		SetCreatorID(team.OperatorID).
		AddAdminIDs(team.OperatorID).
		AddMemberIDs(team.OperatorID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, team model.TeamInputWithID) (*ent.Team, error) {
	save, err := r.client.Team.UpdateOneID(team.OperatorID).SetName(team.Name).SetDescription(team.Description).Save(ctx)
	// TODO: Handle name conflict error
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) SetPrivate(ctx context.Context, private bool, teamID int64) (*ent.Team, error) {
	save, err := r.client.Team.UpdateOneID(teamID).SetPrivate(private).Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) SetAdmins(ctx context.Context, teamID int64, admins []int64) (*model.Reply, error) {
	_, err := r.client.Team.UpdateOneID(teamID).AddAdminIDs(admins...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Updated",
		Data:    nil,
	}, nil
}

func (r *mutationResolver) RemoveAdmins(ctx context.Context, teamID int64, admins []int64) (*model.Reply, error) {
	_, err := r.client.Team.UpdateOneID(teamID).RemoveAdminIDs(admins...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Removed",
		Data:    nil,
	}, nil
}

func (r *mutationResolver) AddMembers(ctx context.Context, teamID int64, members []int64) (*model.Reply, error) {
	_, err := r.client.Team.UpdateOneID(teamID).AddMemberIDs(members...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Added",
		Data:    nil,
	}, nil
}

func (r *mutationResolver) KickMembers(ctx context.Context, teamID int64, members []int64) (*model.Reply, error) {
	_, err := r.client.Team.UpdateOneID(teamID).RemoveMemberIDs(members...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Updated",
		Data:    nil,
	}, nil
}

func (r *queryResolver) GetTeamByName(ctx context.Context, teamName string) (*ent.Team, error) {
	only, err := r.client.Team.Query().Where(team.Name(teamName)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return only, nil
}

func (r *queryResolver) GetTeamByID(ctx context.Context, teamID int64) (*ent.Team, error) {
	get, err := r.client.Team.Get(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return get, nil
}

func (r *queryResolver) GetAvailableTeams(ctx context.Context) ([]*ent.Team, error) {
	all, err := r.client.Team.Query().Where(team.PrivateEQ(false)).All(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (r *queryResolver) TeamMembers(ctx context.Context, teamID int64) ([]*ent.User, error) {
	members, err := r.client.Team.GetX(ctx, teamID).Members(ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *teamResolver) Description(ctx context.Context, obj *ent.Team) (string, error) {
	return obj.Description, nil
}

// Team returns TeamResolver implementation.
func (r *Resolver) Team() TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
