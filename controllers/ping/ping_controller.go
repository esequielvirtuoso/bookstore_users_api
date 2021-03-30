package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// c is a pointer of gin.Context
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "The API is up and running!")
}