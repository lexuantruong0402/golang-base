package grpc

import (
	"smc-wallet-be/src/module/user/services"
)

type UserGRPCServer struct {
	UserService services.IUserService
}

func NewUserGRPCServer() *UserGRPCServer {
	return &UserGRPCServer{}
}
