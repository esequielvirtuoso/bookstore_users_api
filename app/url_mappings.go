package app

import (
	"github.com/esequielvirtuoso/bookstore_users_api/controllers/ping"
	"github.com/esequielvirtuoso/bookstore_users_api/controllers/users"
)

func mapUrls() {
	// This route allows us to test if the service is up.
	// curl -X GET localhost:8080/ping
	router.GET("/ping", ping.Ping)

	// curl -X GET localhost:8080/users/123 -v
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/search", users.SearchUser)

	// curl -X POST localhost:8080/users
	router.POST("/users", users.CreateUser)
}