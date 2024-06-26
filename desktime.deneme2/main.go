package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DeskTimeData struct {
	DesktimeTime int `json:"desktimeTime"`
}

func main() {
	apiKey := "ca6c72efeafd79aca81022a6f00ada05" // Your API key
	userID := "608497"                           // Your user ID

	// Construct the URL without date parameters
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s&id=%s", apiKey, userID)

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Send the GET request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}

	// Unmarshal the JSON data
	var data DeskTimeData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return
	}

	// Print the decoded data
	fmt.Printf("Desktime Time: %d\n", data.DesktimeTime)
}
