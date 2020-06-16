package resource

import "github.com/gin-gonic/gin"

const OK = 200

func HealthCheck(c *gin.Context) {
	c.JSON(OK, gin.H{
		"status": "OK",
	})
}
