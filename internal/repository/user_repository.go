package repository

import (
	"context"

	"github.com/kristiansantos/ms_first/internal/entity"
	"github.com/kristiansantos/ms_first/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UserRepository interface {
	GetAll(filter bson.M) (users entity.Users, err error)
	GetBy(id string) (user *entity.User, err error)
	Create(user entity.User) (entity.User, error)
	Update(id string, user entity.User) (entity.User, error)
	Delete(id string) error
}

type usersRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUsersRepository(db mongodb.Storage, dbCtx context.Context) UserRepository {
	return &usersRepository{
		collection: db.MongoDB.Collection(UsersCollection),
		ctx:        dbCtx,
	}
}

func (r *usersRepository) GetAll(filter bson.M) (users entity.Users, err error) {
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return
	}

	for cursor.Next(r.ctx) {
		document := &entity.User{}
		cursor.Decode(&document)
		users = append(users, document)
	}

	return
}

func (r *usersRepository) GetBy(id string) (user *entity.User, err error) {
	filter := bson.M{"_id": id}

	FindError := r.collection.FindOne(r.ctx, filter).Decode(&user)
	if FindError != nil {
		return &entity.User{}, FindError
	}

	return
}

func (r *usersRepository) Create(user entity.User) (entity.User, error) {
	_, InsertOneError := r.collection.InsertOne(r.ctx, user)
	if InsertOneError != nil {
		return entity.User{}, InsertOneError
	}

	findUser, _ := r.GetBy("filter")

	return *findUser, nil
}

func (r *usersRepository) Update(id string, user entity.User) (entity.User, error) {
	r.collection.UpdateByID(r.ctx, id, user)

	return entity.User{}, nil
}

func (r *usersRepository) Delete(id string) error {
	r.collection.DeleteOne(r.ctx, bson.ObjectIdHex(id))

	return nil
}
