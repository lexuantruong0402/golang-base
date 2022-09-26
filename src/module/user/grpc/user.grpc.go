package grpc_handler

import (
	"context"
	"smc-wallet-be/src/module/user/services"
	"smc-wallet-be/src/proto/pb"

	"gorm.io/gorm"
)

type UserGRPCServer struct {
	UserService services.UserService
}

func NewUserGRPCServer(dbConnection *gorm.DB) *UserGRPCServer {
	return &UserGRPCServer{UserService: *services.NewUserService(dbConnection)}
}

func (s *UserGRPCServer) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user, err := s.UserService.GetUserById(ctx, int(req.GetId()))
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{Id: uint32(user.Id), Username: user.Username, Email: user.Email}, nil
}
