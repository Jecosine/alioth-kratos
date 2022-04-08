package service

import (
	"context"

	pb "github.com/Jecosine/alioth-kratos/api/gql/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BaseService struct {
	pb.UnimplementedBaseServer
}

func NewBaseService() *BaseService {
	return &BaseService{}
}

func (s *BaseService) Ping(ctx context.Context, req *emptypb.Empty) (*pb.PingReply, error) {
	return &pb.PingReply{
		Ping: "Pong",
	}, nil
}
