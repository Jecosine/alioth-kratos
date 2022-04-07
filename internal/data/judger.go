package data

import (
	"context"
	"github.com/Jecosine/alioth-kratos/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type judgerRepo struct {
	data *Data
	log  *log.Helper
}

// NewJudgerRepo .
func NewJudgerRepo(data *Data, logger log.Logger) biz.JudgerRepo {
	return &judgerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *judgerRepo) Save(ctx context.Context, g *biz.Judger) (*biz.Judger, error) {
	return g, nil
}

func (r *judgerRepo) Update(ctx context.Context, g *biz.Judger) (*biz.Judger, error) {
	return g, nil
}

func (r *judgerRepo) FindByID(context.Context, int64) (*biz.Judger, error) {
	return nil, nil
}

func (r *judgerRepo) ListByHello(context.Context, string) ([]*biz.Judger, error) {
	return nil, nil
}

func (r *judgerRepo) ListAll(context.Context) ([]*biz.Judger, error) {
	return nil, nil
}
