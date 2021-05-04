// DTO:
// Data Transfer Object (object that we will transfer between the persistence layer to the application and backward)
package users

import (
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate(changingEmail bool) *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if changingEmail && user.Email == "" {
		return errors.HandleError(errors.BadRequest, "invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.HandleError(errors.BadRequest, "invalid password")
	}
	return nil
}