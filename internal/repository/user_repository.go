package repository

import (
	"github.com/kristiansantos/ms_first/internal/entity"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository interface {
	GetAll(filter bson.M) (users entity.Users, err error)
	GetBy(id string) (user *entity.User, err error)
	Create(user entity.User) (entity.User, error)
	Update(id string, user entity.User) (entity.User, error)
	Delete(id string) error
}
