package dto

type UserCreate struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserUpdate struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
