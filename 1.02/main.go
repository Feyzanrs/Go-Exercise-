package main

import (
	"fmt"
)

func main() {
	helloList := []string{
		"Hello, world",
		"Καλημέρα κόσμε",
		"こんにちは世界",
		"سلام دنیا",
		"Привет, мир",
	}

	// Dizinin eleman sayısını yazdır
	fmt.Println("Length of helloList:", len(helloList))

	// Dizinin son elemanını yazdır
	fmt.Println("Last element of helloList:", helloList[len(helloList)-1])

	// Yanlış indeks erişimini kaldırdım
	// Eğer son elemanı tekrar yazdırmak isterseniz, yukarıdaki gibi yapabilirsiniz
}
