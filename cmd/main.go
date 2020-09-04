package main

import (
	drive "fileshare/pkg/drive"
	"fmt"
	"mime"
	"os"
	filepath "path/filepath"
)

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}

func uploadFile(filepath string, service drive.Service) {
	f, err := os.Open(filepath)

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", err))
	}

	defer f.Close()

	tp := getFilePathMimeType(filepath)

	file, err := drive.CreateFile(service.Drive, filepath, tp, f, "root")
	fmt.Println(file)

	return
}

func main() {
	creditials, err := os.Getwd()
	creditials = creditials + "/credentials"

	service, err := drive.Authorize(creditials)

	if err != nil {
		panic(fmt.Sprintf("Could not create service: %v\n", err))
	}

	uploadFile("C:\\Users\\Giorgi\\Downloads\\IMG_20200811_173954.jpg", service)
}
