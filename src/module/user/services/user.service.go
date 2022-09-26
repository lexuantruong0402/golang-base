package services

import (
	"smc-wallet-be/src/module/user/repositories"
)

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService() *IUserService {
	return &UserService{}
}

func (userService *UserService) CreateUser() {

}
