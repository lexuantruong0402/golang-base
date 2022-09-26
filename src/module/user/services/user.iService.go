package services

import (
	"context"
	"smc-wallet-be/src/module/user/dto"
)

type IUserService interface {
	GetUserById(context.Context, int) (dto.UserDto, error)
}
