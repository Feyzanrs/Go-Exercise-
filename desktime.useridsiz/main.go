package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// DeskTimeData models the JSON structure returned from the API
type DeskTimeData struct {
	DeskTimeTime int `json:"desktimeTime"` // The amount of time in seconds the employee has spent on desktime
}

// User struct to model the JSON structure for user API keys
type User struct {
	Name   string `json:"name"`
	ApiKey string `json:"apiKey"`
}

func fetchDeskTimeData(apiKey string) {
	// Constructing the API URL
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s", apiKey)

	// Making an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// Parsing the JSON response
	var data DeskTimeData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return
	}

	// Calculating the DeskTime in hours and minutes
	hours := data.DeskTimeTime / 3600
	minutes := (data.DeskTimeTime % 3600) / 60

	// Printing the DeskTime value
	fmt.Printf("DeskTime for the employee: %d hours, %d minutes\n", hours, minutes)
}

func main() {
	// Opening JSON file with API keys
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Reading the JSON file
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	// Parsing the JSON data into User slice
	var users []User
	if err := json.Unmarshal(jsonData, &users); err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	// Iterating over users and fetching DeskTime data
	for _, user := range users {
		fmt.Printf("Fetching DeskTime data for %s\n", user.Name)
		fetchDeskTimeData(user.ApiKey)
	}
}
