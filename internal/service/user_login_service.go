package service

import (
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
)

type LogInUserServicer interface {
	Execute(email string, password string) (*entity.User, error)
}
type logInShowService struct {
	repository repository.UserRepository
}

func NewLogInUserService(repository repository.UserRepository) LogInUserServicer {
	return &logInShowService{repository}
}

func (service *logInShowService) Execute(email string, password string) (*entity.User, error) {
	return &entity.User{}, nil
}
