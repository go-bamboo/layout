package service

import (
	"context"

	pb "github.com/go-bamboo/layout/api/layout/admin"
	"github.com/go-bamboo/layout/internal/biz"
)

type AdminService struct {
	pb.UnimplementedQuantBotAdminServer
}

func NewAdminService(uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{}
}

func (s *AdminService) PageBot(ctx context.Context, req *pb.PageBotRequest) (*pb.PageBotReply, error) {
	return nil, nil
}
func (s *AdminService) ListBot(ctx context.Context, req *pb.ListBotRequest) (*pb.ListBotReply, error) {
	return nil, nil
}
func (s *AdminService) GetBot(ctx context.Context, req *pb.GetBotRequest) (*pb.GetBotReply, error) {
	return nil, nil
}

func (s *AdminService) CreateBot(ctx context.Context, req *pb.CreateBotRequest) (*pb.CreateBotReply, error) {
	return nil, nil
}

func (s *AdminService) UpdateBot(ctx context.Context, req *pb.UpdateBotRequest) (*pb.UpdateBotReply, error) {
	return nil, nil
}

func (s *AdminService) DeleteBot(ctx context.Context, req *pb.DeleteBotRequest) (*pb.DeleteBotReply, error) {
	return nil, nil
}
