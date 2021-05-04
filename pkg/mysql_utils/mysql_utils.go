package mysql_utils

import (
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.HandleError(errors.NotFount, "no record matching giving id")
		}
		return errors.HandleError(errors.InternalError, "error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.HandleError(errors.BadRequest, "invalid data")
	}
	return errors.HandleError(errors.InternalError, "error processing request")
}