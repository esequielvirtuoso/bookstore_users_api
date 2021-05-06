// Provide the functionalities or the entry points to interact with users API.
// Take the request, validate if we have all the parameters that we need to
// handle this request and send this handling to the service where we have
// the required business logic.
package users

import (
	"github.com/esequielvirtuoso/bookstore_users_api/domain/users"
	"github.com/esequielvirtuoso/bookstore_users_api/internal/infrastructure/errors"
	"github.com/esequielvirtuoso/bookstore_users_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserID(userIDParam string)(int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil{
		return 0, errors.HandleError(errors.BadRequest,"user id should be a number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true" ))
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.HandleError(errors.BadRequest,"Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true" ))
}

func Update(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.HandleError(errors.BadRequest,"Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true" ))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true" ))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.HandleError(errors.BadRequest, errors.InvalidJsonBody)
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true" ))
}