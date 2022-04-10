package service

import (
	"context"
	"github.com/Jecosine/alioth-kratos/pkg/utils"

	pb "github.com/Jecosine/alioth-kratos/api/auth/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
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
func (s *AuthService) ListAuth(ctx context.Context, req *pb.ListAuthRequest) (*pb.ListAuthReply, error) {
	return &pb.ListAuthReply{}, nil
}
