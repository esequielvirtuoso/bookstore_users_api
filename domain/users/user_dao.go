// DAO
// Data Access Object - The entire logic to persist or read user from/to a given database.
// Access Layer to the database.
package users

import (
	"fmt"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/datasources/mysql/users_db"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_users_api/pkg/logger"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?,?,?,?,?,?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Mock
var (
	usersDB = make(map[int64]*User)
)

func (user *User)Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	return nil
}

func (user *User)Save() * errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	user.Id = userID

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer stmt.Close()

	_, errUpdate := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if errUpdate != nil {
		logger.Error("error when trying to update user", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer stmt.Close()

	if _, errDelete := stmt.Exec(user.Id); errDelete != nil {
		logger.Error("error when trying to delete user by id", err)
		return errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer stmt.Close()

	rows, errFind := stmt.Query(status)
	if errFind != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.HandleError(errors.InternalError, errors.DatabaseError)
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.HandleError(errors.InternalError, errors.DatabaseError)
		}

		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.HandleError(errors.NotFount, fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}