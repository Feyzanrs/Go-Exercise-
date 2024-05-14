package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	printIfNotNil("Count1", count1)
	printIfNotNil("Count2", count2)
	printIfNotNil("Count3", count3)
	printIfNotNil("T", t)
}

func printIfNotNil(name string, ptr interface{}) {
	if ptr != nil {
		fmt.Printf("%s: %#v\n", name, ptr)
	}
}
