package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// DeskTimeData represents the structure for the data we expect from DeskTime API
type DeskTimeData struct {
	DesktimeTime int `json:"desktimeTime"`
}

// fetchDeskTimeData makes an HTTP GET request to DeskTime API and decodes the JSON response.
func fetchDeskTimeData(apiKey, userID, startDate, endDate string) (*DeskTimeData, error) {
	// Construct the URL with query parameters for date range
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s&id=%s&from=%s&to=%s", apiKey, userID, startDate, endDate)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Execute the HTTP request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into DeskTimeData struct
	var data DeskTimeData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// handleDeskTimeEndpoint handles requests to your endpoint
func handleDeskTimeEndpoint(w http.ResponseWriter, r *http.Request) {
	apiKey := "ca6c72efeafd79aca81022a6f00ada05" // Replace with your actual API key
	userID := "608497"                           // Specify the user ID
	startDate := "2024-04-15"                    // Define start date
	endDate := "2024-04-21"                      // Define end date

	data, err := fetchDeskTimeData(apiKey, userID, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to fetch DeskTime data", http.StatusInternalServerError)
		return
	}

	// Convert the DeskTime data to JSON and send it back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/getWeeklyDesktime", handleDeskTimeEndpoint)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
