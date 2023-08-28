package main

import (
	//"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func searchByItem(c *gin.Context) {
	item := c.Param("item")
	book, err := PerformGetRequest(item)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func PerformGetRequest(item string) {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}

	var apiKey = os.Getenv("API_KEY")

	var myurl = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + item + "&location=44.666070,-63.657702&radius=1000&key=" + apiKey

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func main() {
	router := gin.Default()

	//Enable CORS for all origins, methods and headers
	router.Use(corsMiddleware())

	//Defining routes
	router.GET("/search/:item", searchByItem)

	//Run the server
	router.Run("localhost:900")

	//PerformGetRequest()
}
