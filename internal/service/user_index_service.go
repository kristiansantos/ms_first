package service

import (
	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/internal/repository"
	"gopkg.in/mgo.v2/bson"
)

type IndexUserServicer interface {
	GetAllUsers(filter bson.M) (users entity.Users, err error)
}
type userIndexService struct {
	repository repository.UserRepository
}

func NewUserIndexService(repository repository.UserRepository) IndexUserServicer {
	return &userIndexService{repository}
}

func (service *userIndexService) GetAllUsers(filter bson.M) (users entity.Users, err error) {
	return service.repository.GetAll(filter)
}
