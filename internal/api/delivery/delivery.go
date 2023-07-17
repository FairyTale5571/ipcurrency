package delivery

import "github.com/gin-gonic/gin"

type (
	StatusHTTP interface {
		CheckStatus() func(c *gin.Context)
	}

	IPInfoHTTP interface {
		GetIPInfo() func(c *gin.Context)
	}
)
