package service

import (
	"github.com/kristiansantos/ms_first/internal/dto"
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
)

type UpdateUserServicer interface {
	Execute(id string, userDto dto.UserUpdate) (user entity.User, err error)
}
type userUpdateService struct {
	repository repository.UserRepository
}

func NewUserUpdateService(repository repository.UserRepository) UpdateUserServicer {
	return &userUpdateService{repository}
}

func (service *userUpdateService) Execute(id string, userDto dto.UserUpdate) (user entity.User, err error) {
	var userUpdate = entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}

	userUpdate.Populate()

	return service.repository.Update(id, userUpdate)
}
