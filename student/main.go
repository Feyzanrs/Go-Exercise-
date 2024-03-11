package main

import "fmt"

// Student struct'ı bir öğrenciyi temsil eder.
type Student struct {
	Name    string
	Surname string
	Age     int
	Number  int
}

// Kullanıcıdan öğrenci bilgilerini alır ve bir Student struct'ı döndürür.
func GetStudentInfo() Student {
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

// Öğrenci bilgilerini ekrana yazdırma fonksiyonu
func PrintStudentInfo(student Student) {
	fmt.Println("Öğrenci Adı:", student.Name)
	fmt.Println("Öğrenci Soyadı:", student.Surname)
	fmt.Println("Öğrenci Yaşı:", student.Age)
	fmt.Println("Öğrenci Numarası:", student.Number)
}

func main() {

	// Kullanıcı öğrenci bilgilerini al
	student := GetStudentInfo()

	// Öğrenci bilgilerini ekrana yazdır
	PrintStudentInfo(student)

}
