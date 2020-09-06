package main

import (
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"

	"github.com/bakurits/fileshare/pkg/drive"
)

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}

func uploadFile(filepath string, s *drive.Service) {
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("cannot open file: %v", err))
	}

	defer f.Close()

	tp := getFilePathMimeType(filepath)

	file, err := s.CreateFile(filepath, tp, f, "root")
	fmt.Println(file)

	return
}

func main() {
	creditials, err := os.Getwd()
	creditials = creditials + "/credentials"
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}

	service, err := drive.Authorize(creditials)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Enter The File Path \n")
	var filePath string
	if _, err := fmt.Scan(&filePath); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(filePath)

	uploadFile(filePath, service)

}
