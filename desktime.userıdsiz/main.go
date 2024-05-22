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

	// Construct the URL
	url := fmt.Sprintf("https://desktime.com/api/v2/json/company?apiKey=%s", apiKey)

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Send the GET request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body using io.ReadAll instead of ioutil.ReadAll
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
