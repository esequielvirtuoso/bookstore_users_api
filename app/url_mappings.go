package app

import (
	"github.com/esequielvirtuoso/bookstore_users_api/controllers/ping"
	"github.com/esequielvirtuoso/bookstore_users_api/controllers/users"
)

func mapUrls() {
	// Ping
	// This route allows us to test if the service is up.
	// curl -X GET localhost:8080/ping
	router.GET("/ping", ping.Ping)

	// Create User
	// curl -X POST localhost:8080/users
	router.POST("/users", users.Create)

	// Get User
	// curl -X GET localhost:8080/users/123 -v
	router.GET("/users/:user_id", users.Get)
	//router.GET("/search", users.SearchUser)

	// Fully Update User
	//PUT
	router.PUT("/users/:user_id", users.Update)

	// Partially Update User
	router.PATCH("/users/:user_id", users.Update)

	// Delete User
	router.DELETE("/users/:user_id", users.Delete)

	// Search
	router.GET("/internal/users/search", users.Search)

	// Login
	router.POST("/users/login", users.Login)
}