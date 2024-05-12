package main

import (
	"fmt"
)

func main() {
	query, limit, offset := "bat", 10, 20
	query, limit, offset = "ball", 30, 40

	fmt.Println(query, limit, offset)
}
