package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DeskTimeData struct {
	DesktimeTime int `json:"desktimeTime"`
}

func fetchDeskTimeData(apiKey, userID, startDate, endDate string) (*DeskTimeData, error) {
	// Ensure the date is in the correct format "YYYY-MM-DD"
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s&id=%s&from=%s&to=%s", apiKey, userID, startDate, endDate)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data DeskTimeData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func handleDeskTimeEndpoint(w http.ResponseWriter, r *http.Request) {
	apiKey := "ca6c72efeafd79aca81022a6f00ada05"
	userID := "608497"
	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02") // Start date exactly one week ago
	endDate := time.Now().Format("2006-01-02")                     // Today's date

	data, err := fetchDeskTimeData(apiKey, userID, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to fetch DeskTime data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/getWeeklyDesktime", handleDeskTimeEndpoint)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
