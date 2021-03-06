package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
)

func (r *judgeRecordResolver) User(ctx context.Context, obj *ent.JudgeRecord) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateJudgeRecord(ctx context.Context, recordID int64) (*ent.JudgeRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRecordsByProblem(ctx context.Context, problemID int64) ([]*ent.JudgeRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRecordsByUser(ctx context.Context, userID int64) ([]*ent.JudgeRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLastRecordByUserAndProblem(ctx context.Context, problemID int64, userID int64) (*ent.JudgeRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

// JudgeRecord returns JudgeRecordResolver implementation.
func (r *Resolver) JudgeRecord() JudgeRecordResolver { return &judgeRecordResolver{r} }

type judgeRecordResolver struct{ *Resolver }
