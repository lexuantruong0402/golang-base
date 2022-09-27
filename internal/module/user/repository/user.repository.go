package repository

import (
	"smc-wallet-be/internal/module/user/dto"
	"smc-wallet-be/internal/module/user/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindOneById(id int) dto.UserDto {
	var user entities.User

	err := r.db.Where(&entities.User{Id: id}).Take(&user).Error

	if err != nil {
		return dto.UserDto{}
	}

	return dto.UserDto{Id: user.Id, Username: user.Username, Email: user.Email}
}
