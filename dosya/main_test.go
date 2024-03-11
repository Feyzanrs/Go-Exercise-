func TestWriteToFile(t *testing.T) {
	// Mock file
	mockFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatal("Mock dosya oluşturma hatası:", err)
	}
	defer os.Remove(mockFile.Name())

	// Mock input: simulate user input
	mockInput := "test text\n"
	mockScanner := newMockScanner(mockInput)

	// Override standard input with mockScanner
	os.Stdin = mockScanner

	// Call the writeToFile function
	writeToFile(mockFile)

	// Read the content of the mock file
	content, err := os.ReadFile(mockFile.Name())
	if err != nil {
		t.Fatal("Dosya okuma hatası:", err)
	}

	// Expected content
	expectedContent := "test text\n"

	// Compare actual content with expected content
	if string(content) != expectedContent {
		t.Errorf("Beklenen ve alınan içerikler uyuşmuyor.\nBeklenen: %s\nAlınan: %s", expectedContent, string(content))
	}
}

func TestReadFromFile(t *testing.T) {
	// Mock file
	mockFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatal("Mock dosya oluşturma hatası:", err)
	}
	defer os.Remove(mockFile.Name())

	// Write content to the mock file
	content := "test line 1\ntest line 2\n"
	err = os.WriteFile(mockFile.Name(), []byte(content), 0644)
	if err != nil {
		t.Fatal("Dosyaya yazma hatası:", err)
	}

	// Capture standard output
	capturedOutput := captureOutput(func() {
		// Call the readFromFile function
		readFromFile(mockFile)
	})

	// Expected output
	expectedOutput := "Dosya içeriği:\ntest line 1\ntest line 2\n"

	// Compare actual output with expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Beklenen ve alınan çıktılar uyuşmuyor.\nBeklenen:\n%s\nAlınan:\n%s", expectedOutput, capturedOutput)
	}
}

// captureOutput captures the standard output of a function.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf []byte
	r.Read(buf)
	return string(buf)
}
