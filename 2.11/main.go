package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8) // r değeri 0 ile 7 arasında bir tamsayı üretir.
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		fmt.Println(r)
	}
}
