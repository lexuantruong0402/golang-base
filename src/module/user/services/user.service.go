package services

import (
	"context"
	"smc-wallet-be/src/module/user/dto"
	"smc-wallet-be/src/module/user/repository"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(dbConnection *gorm.DB) *UserService {
	return &UserService{UserRepo: *repository.NewUserRepository(dbConnection)}
}

func (userService *UserService) GetUserById(ctx context.Context, id int) (dto.UserDto, error) {
	user := userService.UserRepo.FindOneById(id)
	return user, nil
}
