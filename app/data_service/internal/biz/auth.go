package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Auth is an Auth model.
type Auth struct {
	Hello string
}

// AuthRepo is an Auth repo.
type AuthRepo interface {
	Save(context.Context, *Auth) (*Auth, error)
	Update(context.Context, *Auth) (*Auth, error)
	FindByID(context.Context, int64) (*Auth, error)
	ListByHello(context.Context, string) ([]*Auth, error)
	ListAll(context.Context) ([]*Auth, error)
}

// AuthUsecase is an Auth usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

// NewAuthUsecase new an Auth usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAuth creates an Auth, and returns the new Auth.
func (uc *AuthUsecase) CreateAuth(ctx context.Context, a *Auth) (*Auth, error) {
	uc.log.WithContext(ctx).Infof("CreateAuth: %v", a.Hello)
	return uc.repo.Save(ctx, a)
}

func (uc *AuthUsecase) GetRole(ctx context.Context, role string) (string, error) {
	uc.log.WithContext(ctx).Infof("Get Role: %v", role)
	return role, nil
}
