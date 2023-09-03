package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type location struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Hours       string `json:"hours"`
	Coordinates string `json:"coordinates"`
}

func searchByItem(c *gin.Context) {
	item := c.Param("item")
	result := PerformGetRequest(item)

	if result == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func PerformGetRequest(item string) string {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
		fmt.Println("Problem retreiving API key")
		return ""
	}

	var apiKey = os.Getenv("API_KEY")
	var myurl = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + item + "&location=44.666070,-63.657702&radius=1000&key=" + apiKey

	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println("Problem with executing Google maps API")
		return ""
	}

	defer response.Body.Close()

	//fmt.Println("Status code: ", response.StatusCode)

	content, _ := io.ReadAll(response.Body)

	//fmt.Println(string(content))
	result := string(content)
	addressFinderValue := content
	addressFinder(addressFinderValue)

	return result
}

func addressFinder(content []byte) {
	checkValid := json.Valid(content)

	if checkValid {
		fmt.Println("json format valid")
		var data map[string]interface{}
		json.Unmarshal(content, &data)

		fmt.Println(data["results"].(string))
	} else {
		fmt.Println("json format invlaid")
	}
}

func main() {
	router := gin.Default()

	//Enable CORS for all origins, methods and headers
	router.Use(corsMiddleware())

	//Defining routes
	router.GET("/search/:item", searchByItem)

	//Run the server
	router.Run("localhost:9000")

	//PerformGetRequest()
}
