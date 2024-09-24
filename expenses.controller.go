package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type expense struct {
	ID       string  `json: "id"`
	Title    string  `json: "title"`
	Amount   float64 `json: "amount"`
	DateTime string  `json: "date"`
}

var expenses = []expense{
	{
		ID:       "1",
		Title:    "Groceries",
		Amount:   150.75,
		DateTime: time.Now().Format("2006-01-02 15:04:05"), // Using current date-time
	},
	{
		ID:       "2",
		Title:    "Electricity Bill",
		Amount:   60.25,
		DateTime: "2024-09-01 10:30:00",
	},
	{
		ID:       "3",
		Title:    "Internet Subscription",
		Amount:   40.00,
		DateTime: "2024-09-10 14:15:00",
	},
}

func getExpenses(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, expenses)
}

// postAlbums adds an album from JSON received in the request body.
func createExpense(context *gin.Context) {
	var newExpense expense

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := context.BindJSON(&newExpense); err != nil {
		return
	}

	// Add the new album to the slice.
	expenses = append(expenses, newExpense)
	context.IndentedJSON(http.StatusCreated, newExpense)
}

func getExpensesByID(context *gin.Context) {
	id := context.Param("id")

	for _, e := range expenses {
		if e.ID == id {
			context.IndentedJSON(http.StatusOK, e)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense not found"})
}
