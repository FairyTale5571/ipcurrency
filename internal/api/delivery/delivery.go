package delivery

import "github.com/gin-gonic/gin"

type (
	IPInfoHTTP interface {
		GetIPInfo() func(c *gin.Context)
	}
)
