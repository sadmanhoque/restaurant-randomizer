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

type location struct {
	Name         string `json:"name"`
	StoreAddress string `json:"storeAddress"`
	Error        bool   `json:"error"`
}

func searchByItem(c *gin.Context) {
	item := c.Param("item")
	result := PerformGetRequest(item)

	if result.Error {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		fmt.Println(result.Name)
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

func PerformGetRequest(item string) location {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
		fmt.Println("Problem retreiving API key")
		var result location
		result.Name = "Problem retreiving API key"
		result.Error = true
		return result
	}

	var apiKey = os.Getenv("API_KEY")
	var myurl = "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + item + "&location=44.647646,-63.590651&radius=1000&key=" + apiKey

	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println("Problem with executing Google maps API")
		var result location
		result.Name = "Problem with executing Google maps API"
		result.Error = true
		return result
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

func addressFinder(jsonData []byte) location {
	checkValid := json.Valid(jsonData)
	var output location

	if checkValid {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			fmt.Println("Error:", err)
			output.Name = "error in unmarshaling json data from google"
			output.Error = true
			return output
		}

		results, found := data["results"].([]interface{})
		if !found {
			//fmt.Println("Results not found in JSON")
			output.Name = "Nothing found :("
			output.StoreAddress = "N/A"
			output.Error = false
			return output
		}

		result := results[rand.Intn(len(results)-0)]

		resultMap, isMap := result.(map[string]interface{})
		if !isMap {
			//fmt.Println("Invalid result format")
			output.Name = "error: unexpected json format"
			output.Error = true
			return output
		}

		//name, nameFound := resultMap["name"].(string)
		//address := resultMap["formatted_address"].(string)

		output.Name = resultMap["name"].(string)
		output.StoreAddress = resultMap["formatted_address"].(string)

		//fmt.Println(lat)

		fmt.Println("Name:", output.Name)
		output.Error = false
		return output

	} else {
		//fmt.Println("json format invlaid")
		output.Name = "error"
		return output
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
