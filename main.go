package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/expenses", getExpenses)
	router.POST("/expenses", createExpense)
	router.GET("/expenses/:id", getExpensesByID)

	router.Run("localhost:8090")
}
