package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	fmt.Println(c.Request.Header)
	c.Next()
}
