package terminal

import (
	"errors"
	"fmt"
	"github.com/bakurits/fileshare/pkg/drive"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type Terminal struct {
	drive    *drive.Service
	drivePWD string
}

func GetInstance() *Terminal {
	return &Terminal{
		drive:    nil,
		drivePWD: "",
	}
}

func (t *Terminal) Execute(cmdString string) error {
	cmdString = strings.TrimSuffix(cmdString, "\n")
	var command []string = strings.Fields(cmdString)

	if len(command) < 1 {
		return errors.New("empty command")
	}

	switch command[0] {
	case "exit":
		// just exit from program
		os.Exit(0)
	case "authorize":
		// only authorize and save service variable in Terminal struct
		return t.authorize(command)
	case "createfile":
		return t.processCreateFile(command)
	}

	return nil
}

func (t *Terminal) authorize(command []string) error {
	if len(command) < 2 {
		return errors.New("authorize: credentials have not passed")
	}
	service, err := drive.Authorize(command[1])
	t.drive = service

	return err
}

func (t *Terminal) processCreateFile(command []string) error {
	if len(command) < 2 {
		return errors.New("createfile: arguments not passed")
	}
	if t.drive == nil {
		return errors.New("createfile: you are not authorized")
	}

	err := uploadFile(command[1], t.drive)

	return err
}

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}

func uploadFile(filepath string, s *drive.Service) error {
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	tp := getFilePathMimeType(filepath)

	_, err = s.CreateFile(filepath, tp, f, "root")

	return err
}
