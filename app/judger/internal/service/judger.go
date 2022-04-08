package service

import (
	"context"
	"github.com/Jecosine/alioth-kratos/app/judger/internal/biz"

	pb "github.com/Jecosine/alioth-kratos/api/judger/v1"
)

type JudgerService struct {
	pb.UnimplementedJudgerServer
	uc *biz.JudgerUsecase
}

func NewJudgerService(uc *biz.JudgerUsecase) *JudgerService {
	return &JudgerService{
		uc: uc,
	}
}

func (s *JudgerService) Ping(ctx context.Context, req *pb.PingJudgerRequest) (*pb.PingJudgerReply, error) {
	return &pb.PingJudgerReply{
		Id: 0,
		Data: &pb.JudgerStatus{
			Status:             4,
			CurrentTasksAmount: 10,
			EstimateTime:       1.5,
		},
	}, nil
}
func (s *JudgerService) Submit(ctx context.Context, req *pb.SubmitJudgerRequest) (*pb.SubmitJudgerReply, error) {
	judgeRequest := req.Payload
	judgeReply, err := s.uc.ProcessRequest(ctx, judgeRequest)
	if err != nil {
		return nil, err
	}
	return judgeReply, nil
}
