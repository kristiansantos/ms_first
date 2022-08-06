package service

import (
	"github.com/kristiansantos/ms_first/internal/dto"
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
)

type CreateUserServicer interface {
	Execute(userDto dto.UserCreate) (user entity.User, err error)
}
type userCreateService struct {
	repository repository.UserRepository
}

func NewUserCreateService(repository repository.UserRepository) CreateUserServicer {
	return &userCreateService{repository}
}

func (service *userCreateService) Execute(userDto dto.UserCreate) (user entity.User, err error) {
	var userCreate = entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}

	userCreate.Populate()

	return service.repository.Create(userCreate)
}
