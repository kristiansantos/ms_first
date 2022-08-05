package service

import (
	"github.com/kristiansantos/ms_first/internal/repository"
)

type DeleteUserServicer interface {
	DeleteUser(id string) error
}
type userDeleteService struct {
	repository repository.UserRepository
}

func NewUserDeleteService(repository repository.UserRepository) DeleteUserServicer {
	return &userDeleteService{repository}
}

func (service *userDeleteService) DeleteUser(id string) (err error) {
	return service.repository.Delete(id)
}
