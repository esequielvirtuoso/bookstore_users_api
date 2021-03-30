// The entire business logic must be here
package services

import (
	"github.com/esequielvirtuoso/bookstore_users_api/domain/users"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr){
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	result := &users.User{Id: userId}
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