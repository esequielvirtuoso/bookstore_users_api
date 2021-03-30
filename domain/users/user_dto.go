// DTO:
// Data Transfer Object (object that we will transfer between the persistence layer to the application and backward)
package users

import (
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"strings"
)

type User struct {
	Id          int64	`json:"id"`
	FirstName   string	`json:"first_name"`
	LastName    string	`json:"last_name"`
	Email       string	`json:"email"`
	DateCreated string	`json:"date_created"`
}

func (user *User)Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}