package main

import "fmt"
import "github.com/gin-gonic/gin"
import "expenser-api/resource"

func main() {
	fmt.Println("Welcome to Expenser.")

	router := gin.Default()
	router.GET("/health", resource.HealthCheck)
	router.POST("/expenses", resource.AddExpense)

	err := router.Run(":8000")
	fmt.Println(err)
}
