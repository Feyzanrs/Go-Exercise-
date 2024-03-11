package main

import (
	"fmt"
)

func ToplaFonksiyonu(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}

	return toplam, nil
}

func main() {
	// Test için kullanılacak örnek giriş verisi
	sayilar := []int{1, 2, 3, 4, 5, 6}

	// ToplaFonksiyonu fonksiyonunu çağırıyoruz
	toplam, err := ToplaFonksiyonu(sayilar)

	// Hata kontrolü
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Sonucu ekrana yazdır
	fmt.Println("Toplam:", toplam)
}
