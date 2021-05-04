package app

import (
	"github.com/esequielvirtuoso/bookstore_users_api/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	// NOTE: This is the only layer we are defining and using the HTTP server
	router = gin.Default()
)

func StartApplication() {
	mapUrls() // defining maps
	logger.Info("about to start the application...")
	router.Run(":8080") // running on 8080 port
}