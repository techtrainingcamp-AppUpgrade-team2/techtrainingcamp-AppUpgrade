package service

import (
	"github.com/gin-gonic/gin"
)

func Demo(c *gin.Context) {
	c.String(200, "Demo")
}
