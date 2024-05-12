package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestÇarpma(t *testing.T) {
	// Testify paketini kullanarak assertion özelliklerini kullanıyoruz.
	assert := assert.New(t)

	// Test için kullanılacak örnek giriş verisi
	sayilar := []int{2, 3, 4, 5, 6, 7, 8, 9}

	// Çarpma fonksiyonunu çağırıyoruz
	çarpım, err := ÇarpmaFonksiyonu(sayilar)

	// Assertion kullanarak beklenen değerleri kontrol ediyoruz.
	assert.Nil(err)           // Hata olmamalı
	assert.Equal(720, çarpım) // Çarpım 720 olmalı
}
