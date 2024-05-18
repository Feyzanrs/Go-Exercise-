package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
