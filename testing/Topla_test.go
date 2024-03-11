package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopla(t *testing.T) {
	// Testify paketini kullanarak assertion özelliklerini kullanıyoruz.
	assert := assert.New(t)

	// Test için kullanılacak örnek giriş verisi
	sayilar := []int{1, 2, 3, 4, 5, 6}

	// Topla fonksiyonunu çağırıyoruz
	toplam, err := ToplaFonksiyonu(sayilar)

	// Assertion kullanarak beklenen değerleri kontrol ediyoruz.
	assert.Nil(err)          // Hata olmamalı
	assert.Equal(21, toplam) // Toplam 21 olmalı
}
