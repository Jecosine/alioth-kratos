package biz

import (
	v1 "github.com/Jecosine/alioth-kratos/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	"context"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Judger is a Judger model.
type Judger struct {
	Hello string
}

// JudgerRepo is a Judger repo.
type JudgerRepo interface {
	Save(context.Context, *Judger) (*Judger, error)
	Update(context.Context, *Judger) (*Judger, error)
	FindByID(context.Context, int64) (*Judger, error)
	ListByHello(context.Context, string) ([]*Judger, error)
	ListAll(context.Context) ([]*Judger, error)
}

// JudgerUsecase is a Judger usecase.
type JudgerUsecase struct {
	repo JudgerRepo
	log  *log.Helper
}

// NewJudgerUsecase new a Judger usecase.
func NewJudgerUsecase(repo JudgerRepo, logger log.Logger) *JudgerUsecase {
	return &JudgerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateJudger creates a Judger, and returns the new Judger.
func (uc *JudgerUsecase) CreateJudger(ctx context.Context, g *Judger) (*Judger, error) {
	uc.log.WithContext(ctx).Infof("CreateJudger: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
