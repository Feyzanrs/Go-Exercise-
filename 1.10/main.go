package main

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)
	count++
	count--
	count -= 5
	fmt.Println(count)
	name := "John"
	name += " Simith"
	fmt.Println("Hello,", name)
}
