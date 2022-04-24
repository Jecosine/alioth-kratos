package data

import (
	"context"

	"github.com/Jecosine/alioth-kratos/app/data_service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) Save(ctx context.Context, g *biz.Auth) (*biz.Auth, error) {
	return g, nil
}

func (r *authRepo) Update(ctx context.Context, g *biz.Auth) (*biz.Auth, error) {
	return g, nil
}

func (r *authRepo) FindByID(context.Context, int64) (*biz.Auth, error) {
	return nil, nil
}

func (r *authRepo) ListByHello(context.Context, string) ([]*biz.Auth, error) {
	return nil, nil
}

func (r *authRepo) ListAll(context.Context) ([]*biz.Auth, error) {
	return nil, nil
}
