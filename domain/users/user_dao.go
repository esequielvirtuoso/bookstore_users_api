// DAO
// Data Access Object - The entire logic to persist or read user from/to a given database.
// Access Layer to the database.
package users

import (
	"fmt"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
)

// Mock
var (
	usersDB = make(map[int64]*User)
)

func (user *User)Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User)Save() * errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}