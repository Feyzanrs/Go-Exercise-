package main

import "fmt"

func main() {
	// Bir dizi oluşturun
	numbers := []int{72, 90, 78, 57, 46}

	var total int
	count := 0

	for _, num := range numbers {
		total += num
		count++
	}

	average := float64(total) / float64(count) // Ortalamayı ondalık sayı olarak hesaplayın

	fmt.Printf("series average: %.2f\n", average) // Ortalamayı ekrana yazdırın
}
