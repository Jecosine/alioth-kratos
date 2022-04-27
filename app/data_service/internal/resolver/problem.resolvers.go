package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/model"
)

func (r *mutationResolver) AddProblem(ctx context.Context, problem model.ProblemInput) (*ent.Problem, error) {
	save, err := r.client.Problem.Create().SetTitle(problem.Title).SetAuthorID(problem.Author).SetContent(problem.Content).SetCreatedTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) ModifyProblem(ctx context.Context, problem model.ProblemInputWithID) (*ent.Problem, error) {
	save, err := r.client.Problem.UpdateOneID(problem.ID).SetTitle(problem.Title).SetContent(problem.Content).Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (r *mutationResolver) DeleteProblem(ctx context.Context, problemID int64) (*model.Reply, error) {
	err := r.client.Problem.DeleteOneID(problemID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		Code:    200,
		Message: "Deleted",
		Data:    nil,
	}, nil
}

func (r *queryResolver) GetAllProblems(ctx context.Context) ([]*ent.Problem, error) {
	all, err := r.client.Problem.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (r *queryResolver) GetProblemsWithFilter(ctx context.Context, filter *model.ProblemFilter) ([]*ent.Problem, error) {
	//r.client.Problem.Query().Where(	)
	return nil, nil
}
