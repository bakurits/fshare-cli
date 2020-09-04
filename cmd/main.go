package main

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/drive"
	"mime"
	"os"
	"path/filepath"
)

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}

func uploadFile(filepath string, service drive.Service) {
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
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
	if err != nil {
		fmt.Sprintf("cannot open file: %v", err)
	}

	service, err := drive.Authorize(creditials)
	uploadFile("C:\\Users\\Giorgi\\Downloads\\IMG_20200811_164530.jpg", service)
}
