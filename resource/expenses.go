package resource

import (
	"expenser-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense service.Expense
	e := c.BindJSON(&expense)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, "failed to read expense request body.")
	}

	go service.AddExpense(expense)

	c.JSON(201, gin.H{})
}
