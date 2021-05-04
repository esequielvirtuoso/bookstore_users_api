// The entire business logic must be here
package services

import (
	"github.com/esequielvirtuoso/bookstore_users_api/domain/users"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_users_api/pkg/crypto_utils"
	"github.com/esequielvirtuoso/bookstore_users_api/pkg/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUsers(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr){
	if userId <= 0 {
		return nil, errors.HandleError("bad request","invalid user id")
	}

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(true); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if err := user.Validate(user.Email != ""); err != nil {
			return nil, err
		}
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		if err := user.Validate(true); err != nil {
			return nil, err
		}
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}


	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUsers(status string) (users.Users, *errors.RestErr){
	dao := &users.User{}
	return dao.FindByStatus(status)
}