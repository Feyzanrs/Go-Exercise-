package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a Student struct
type Student struct {
	Name    string
	Surname string
	Age     int
	Number  int
}

// Function to print student information
func PrintStudentInfo(student Student) {
	fmt.Printf("Name: %s\nSurname: %s\nAge: %d\nNumber: %d\n", student.Name, student.Surname, student.Age, student.Number)
}

func TestPrintStudentInfo(t *testing.T) {
	// Test için örnek bir öğrenci oluştur
	student := Student{
		Name:    "John",
		Surname: "Doe",
		Age:     25,
		Number:  12345,
	}

	// Student struct'ını yazdıran fonksiyonu çağır
	PrintStudentInfo(student)
}

func TestGetStudentInfo(t *testing.T) {
	// Test için örnek bir öğrenci bilgisi oluştur
	fakeInput := "Jane\nDoe\n22\n67890\n"

	// Öğrenci bilgilerini kullanıcının girdiği gibi taklit etmek için
	// bir yapay giriş oluştur
	var inputFunc = func() string {
		return fakeInput
	}

	// Test için öğrenci bilgilerini al
	student := GetStudentInfo(inputFunc)

	// Assertion kullanarak beklenen değerleri kontrol et
	assert.Equal(t, "Jane", student.Name)
	assert.Equal(t, "Doe", student.Surname)
	assert.Equal(t, 22, student.Age)
	assert.Equal(t, 67890, student.Number)
}

// Kullanıcıdan öğrenci bilgilerini taklit edilmiş bir girişle alır ve bir Student struct'ı döndürür
func GetStudentInfo(inputFunc func() string) Student {
	var ad, soyad string
	var yas, numara int

	fmt.Print("Öğrenci Adı:")
	fmt.Scanln(&ad)

	fmt.Print("Öğrenci Soyadı:")
	fmt.Scanln(&soyad)

	fmt.Print("Öğrenci Yaşı:")
	fmt.Scanln(&yas)

	fmt.Print("Öğrenci Numarası:")
	fmt.Scanln(&numara)

	// Yeni bir student struct'ı oluşturup değerleri atama
	student := Student{
		Name:    ad,
		Surname: soyad,
		Age:     yas,
		Number:  numara,
	}

	return student
}
