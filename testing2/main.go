package main

import (
	"fmt"
)

func ÇarpmaFonksiyonu(sayilar []int) (int, error) {
	çarpim := 1
	for i := range sayilar {
		çarpim = çarpim * sayilar[i]
	}

	return çarpim, nil
}

func main() {
	//Test için kullanılacak örnek giriş verisi
	sayilar := []int{1, 2, 3, 4, 5, 6}

	//ÇarpmaFonksiyonu fonksiyonunu çağıırıyoruz
	çarpim, err := ÇarpmaFonksiyonu(sayilar)

	//Hata kontrolü
	if err != nil {
		fmt.Print("Hata:", err)
		return
	}

	//Sonucu ekrana yazdır
	fmt.Println("Çarpim:", çarpim)

}
