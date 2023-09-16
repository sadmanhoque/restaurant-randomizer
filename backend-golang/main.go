package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type location struct {
// 	Name        string `json:"name"`
// 	Address     string `json:"address"`
// }

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
	//result := string(content)
	addressFinderValue := content
	result := addressFinder(addressFinderValue)

	return result
}

func addressFinder(jsonData []byte) string {
	checkValid := json.Valid(jsonData)

	if checkValid {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			fmt.Println("Error:", err)
			return "error"
		}

		results, found := data["results"].([]interface{})
		if !found {
			fmt.Println("Results not found in JSON")
			return "error"
		}

		result := results[rand.Intn(len(results)-0)]

		resultMap, isMap := result.(map[string]interface{})
		if !isMap {
			fmt.Println("Invalid result format")
			return "error"
		}

		name, nameFound := resultMap["name"].(string)

		if nameFound {
			fmt.Println("Name:", name)
			return name
		}

		return "error"

	} else {
		fmt.Println("json format invlaid")
		return "error"
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
