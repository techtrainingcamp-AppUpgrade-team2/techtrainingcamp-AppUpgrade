package service

import (
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.String(200, "Pong")
}
