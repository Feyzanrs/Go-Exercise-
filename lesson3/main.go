package main

import "fmt"

func main() {

	total := 0

	for number := 1; number <= 100; number++ {
		total += number
	}

	fmt.Println("the sum of numbers from 1 to 100 :", total)
}
