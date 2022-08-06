package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created"`
	UpdatedAt time.Time `bson:"updated"`
}

type Users []*User

func NewUser() *User {
	return &User{}
}

func (u *User) Populate() {
	if u.Id == "" {
		u.Id = uuid.New().String()
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	} else {
		u.UpdatedAt = time.Now()
	}
}
