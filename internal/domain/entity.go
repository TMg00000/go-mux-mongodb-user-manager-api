package domain

import (
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id           bson.ObjectID `bson:"_id,omitempty"`
	Name         string
	Email        string
	Password     string `json:"-"`
	PasswordHash string
}

func NewLoginUser(email, password string) *User {
	if (email == "" && !strings.Contains(email, "@")) || password == "" {
		return nil
	}

	if len(password) < 8 {
		return nil
	}

	return &User{
		Email:    email,
		Password: password,
	}
}

func NewUser(name, email, password string) *User {
	if name == "" || email == "" || password == "" {
		return nil
	}

	if len(name) < 3 || len(name) > 50 {
		return nil
	}

	if len(password) < 8 {
		return nil
	}

	if !strings.Contains(email, "@") {
		return nil
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
