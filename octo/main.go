package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanf("%f", &num1)
	if err != nil {
		fmt.Println("Error reading the first number:", err)
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanf("%f", &num2)
	if err != nil {
		fmt.Println("Error reading the second number:", err)
		return
	}

	if num1 > num2 {
		fmt.Printf("%.2f is greater than %.2f\n", num1, num2)
	} else if num2 > num1 {
		fmt.Printf("%.2f is greater than %.2f\n", num2, num1)
	} else {
		fmt.Println("Both numbers are equal.")
	}
}
