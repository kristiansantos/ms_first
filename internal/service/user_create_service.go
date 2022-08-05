package service

import (
	"github.com/kristiansantos/ms_first/internal/dto"
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
)

type CreateUserServicer interface {
	CreateUser(userDto dto.UserCreate) (user entity.User, err error)
}
type userCreateService struct {
	repository repository.UserRepository
}

func NewUserCreateService(repository repository.UserRepository) CreateUserServicer {
	return &userCreateService{repository}
}

func (service *userCreateService) CreateUser(userDto dto.UserCreate) (user entity.User, err error) {
	var userCreate = entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}

	return service.repository.Create(userCreate)
}
