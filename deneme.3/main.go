package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const apiKey = "ca6c72efeafd79aca81022a6f00ada05"

type EmployeeResponse struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TimeIn string `json:"time_in"`
}

func main() {
	url := fmt.Sprintf("https://desktime.com/api/v2/json/employee?apiKey=%s", apiKey)

	client := resty.New()

	resp, err := client.R().
		SetResult(&EmployeeResponse{}).
		Get(url)

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	employeeResponse := resp.Result().(*EmployeeResponse)

	for _, employee := range employeeResponse.Employees {
		fmt.Printf("Employee ID: %d\n", employee.ID)
		fmt.Printf("Employee Name: %s\n", employee.Name)
		fmt.Printf("Time In: %s\n", employee.TimeIn)
		fmt.Println("----------------------------")
	}
}
