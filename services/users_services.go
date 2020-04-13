package services

import (
	"github.com/macrusal/bookstore_users-api/datasources/mysql/users_db"
	"github.com/macrusal/bookstore_users-api/domain/users"
	"github.com/macrusal/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := &users.User{Id:userId}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}