package service

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/biz"
	"github.com/Jecosine/alioth-kratos/app/data_service/pkg/model"
	"github.com/Jecosine/alioth-kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/errors"

	pb "github.com/Jecosine/alioth-kratos/api/auth/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc *biz.AuthUsecase
}

func NewAuthService(uc *biz.AuthUsecase) *AuthService {
	return &AuthService{
		uc: uc,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *AuthService) CurrentUser(ctx context.Context, req *emptypb.Empty) (*pb.CurrentUserReply, error) {
	claims, err := utils.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CurrentUserReply{
		User:    claims.User,
		Expired: claims.Expired,
	}, nil
}
func (s *AuthService) HasRole() func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.RoleType, id *string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.RoleType, id *string) (res interface{}, err error) {
		fmt.Print("called")
		r, err := s.uc.GetRole(ctx, "USER")

		if err != nil {
			return nil, errors.InternalServer("Cannot get role", "Failed")
		}
		fmt.Printf("role in service %s", r)
		if role.String() != r {
			return nil, errors.Forbidden("Forbidden", "no privilege")
		}
		return next(ctx)
	}
}
func (s *AuthService) ListAuth(ctx context.Context, req *pb.ListAuthRequest) (*pb.ListAuthReply, error) {
	return &pb.ListAuthReply{}, nil
}
