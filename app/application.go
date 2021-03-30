package app

import (
	"github.com/gin-gonic/gin"
)

var (
	// NOTE: This is the only layer we are defining and using the HTTP server
	router = gin.Default()
)

func StartApplication() {
	mapUrls() // defining maps
	router.Run(":8080") // running on 8080 port
}