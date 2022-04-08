package biz

import (
	v1 "github.com/Jecosine/alioth-kratos/api/judger/v1"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	"context"
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
	ProcessRequest(context.Context, *proto.JudgeRequestProto) (*v1.SubmitJudgerReply, error)
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
func (uc *JudgerUsecase) CreateJudger(ctx context.Context, j *Judger) (*Judger, error) {
	uc.log.WithContext(ctx).Infof("CreateJudger: %v", j.Hello)
	return uc.repo.Save(ctx, j)
}

func (uc *JudgerUsecase) ProcessRequest(ctx context.Context, judgeRequest *proto.JudgeRequestProto) (*v1.SubmitJudgerReply, error) {
	uc.log.Infof("Received judge request: %v", judgeRequest)
	judgeReply, err := uc.repo.ProcessRequest(ctx, judgeRequest)
	if err != nil {
		return nil, err
	}
	return judgeReply, err
}
