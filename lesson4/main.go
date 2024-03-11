package main

import "fmt"

func main() {
	for sayi := 1; sayi <= 100; sayi++ {
		if sayi%2 == 1 {
			fmt.Println(sayi)
		}
	}
}
