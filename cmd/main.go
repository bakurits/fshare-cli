package main

import (
	"bufio"
	"fmt"
	"github.com/bakurits/fileshare/pkg/terminal"
	"log"
	"os"
)

func main() {
	//creditials, err := os.Getwd()

	terminal := terminal.GetInstance()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		err = terminal.Execute(cmdString)
		if err != nil {
			log.Println(err)
		}
	}

}
