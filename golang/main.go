package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book one", Author: "Author one", Quantity: 3},
	{ID: "2", Title: "Book two", Author: "Author two", Quantity: 5},
	{ID: "3", Title: "Book three", Author: "Author three", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:5030")
}
