package main

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/drive"
	"os"
)

func main() {
	creditials, err := os.Getwd()
	creditials = creditials + "/credentials"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(creditials)

	client, err := auth.GetHTTPClient(creditials)

	if err != nil {
		fmt.Println(err)
	}

	service, err := drive.NewService(client)

	if err != nil {
		panic(fmt.Sprintf("Could not create service: %v\n", err))
	}

	dir, err := drive.CreateDir(service.Drive, "My Folder", "root")

	if err != nil {
		panic(fmt.Sprintf("Could not create dir: %v\n", err))
	}

	fmt.Println("kargadaa yvelaferi")

	fmt.Println(dir)
}
