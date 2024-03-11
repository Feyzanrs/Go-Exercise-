package main

import "fmt"

func main() {
	// Define an example array
	series := []int{82, 25, 17, 93, 60, 3}

	// Assume first element as smallest
	smallest := series[0]

	// Find the smallest element by rotating through the array
	for _, element := range series {
		if element < smallest {
			smallest = element
		}
	}

	// Print smallest element to screen
	fmt.Println("Smallest element in the array:", smallest)
}
