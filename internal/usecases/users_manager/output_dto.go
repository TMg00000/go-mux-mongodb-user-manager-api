package users_manager

import "go.mongodb.org/mongo-driver/v2/bson"

type UserCreateOutput struct {
	Name  string `bson:"name" `
	Email string `bson:"email" `
}

type UserGetAllOutput struct {
	Id           bson.ObjectID `bson:"_id,omitempty"`
	Name         string
	Email        string
	PasswordHash string
}
