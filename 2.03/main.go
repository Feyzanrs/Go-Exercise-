package main

import "fmt"

func main() {
	input := 6

	if input < 0 {
		fmt.Println("input can't be a negative")

	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}

}
