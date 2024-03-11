package main

import "fmt"

func main() {
	var number int

	fmt.Print("enter a number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(number, "an ever number.")
	} else {
		fmt.Println(number, "a single number.")
	}
}
