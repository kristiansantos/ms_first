package service

import (
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
)

type ShowUserServicer interface {
	Execute(id string) (*entity.User, error)
}
type userShowService struct {
	repository repository.UserRepository
}

func NewUserShowService(repository repository.UserRepository) ShowUserServicer {
	return &userShowService{repository}
}

func (service *userShowService) Execute(id string) (*entity.User, error) {
	return service.repository.GetBy(id)
}
