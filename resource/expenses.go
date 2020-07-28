package resource

import (
	"expenser-api/dto"
	"expenser-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetExpenses(c *gin.Context) {
	expenses := service.GetExpenses()
	c.JSON(200, expenses)
}

func UpdateExpense(c *gin.Context) {
	var expense dto.Expense
	e := c.BindJSON(&expense)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, "failed to read expense request body.")
	}

	service.UpdateExpense(expense)

	c.JSON(204, gin.H{})
}

func AddExpense(c *gin.Context) {
	var expense dto.Expense
	e := c.BindJSON(&expense)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, "failed to read expense request body.")
	}

	service.AddExpense(expense)

	c.JSON(201, gin.H{})
}
