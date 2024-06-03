package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeskTimeData API'den dönen JSON yapısını modellemek için kullanılır
type DeskTimeData struct {
	DeskTimeTime int `json:"desktimeTime"`
}

func fetchDeskTimeData(apiKey string) {
	// API URL'ini oluşturma
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s", apiKey)

	// HTTP GET isteği
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	// Yanıtı okuma
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// JSON yanıtını parse etme
	var data DeskTimeData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return
	}

	// DeskTime süresini saat ve dakika olarak hesaplama
	hours := data.DeskTimeTime / 3600
	minutes := (data.DeskTimeTime % 3600) / 60

	// DeskTime değerini yazdırma
	fmt.Printf("DeskTime for the employee: %d hours, %d minutes\n", hours, minutes)
}

func main() {
	// API anahtarı
	apiKey := "ca6c72efeafd79aca81022a6f00ada05"

	// DeskTime verilerini getir
	fetchDeskTimeData(apiKey)
}
