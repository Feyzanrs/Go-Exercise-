package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Dosya İşleme Uygulaması")

	var fileName string

	// Dosya adını kullanıcıdan al
	fmt.Print("Dosya adını girin: ")
	fmt.Scanln(&fileName)

	// Dosyayı aç veya oluştur
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Dosya açma veya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	for {
		fmt.Println("1. Dosyaya Yaz")
		fmt.Println("2. Dosyadan Oku")
		fmt.Println("3. Çıkış")

		var choice int
		fmt.Print("Seçiminizi yapın (1-3): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			writeToFile(file)
		case 2:
			readFromFile(file)
		case 3:
			fmt.Println("Çıkış yapılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim. Lütfen tekrar deneyin.")
		}
	}
}

func writeToFile(file *os.File) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Dosyaya yazılacak metni girin: ")
	if scanner.Scan() {
		text := scanner.Text()
		_, err := file.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Dosyaya yazma hatası:", err)
		} else {
			fmt.Println("Dosyaya yazma başarılı.")
		}
	}
}

func readFromFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	fmt.Println("Dosya içeriği:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Dosya okuma hatası:", err)
	}
}
