package drivemanager

import (
	"encoding/json"
	"errors"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// NewCreateDirCommand : generates of command createdir
func NewCreateDirCommand() *cobra.Command {

	// createdirCmd represents the createdir command
	var createdirCmd = &cobra.Command{
		Use:   "createdir",
		Short: "creation of directory in google drive",
		Long:  `creation of directory in google drive, first argument is dir you want to create, second argument is parent directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateDir(args)
		},
	}
	return createdirCmd
}

func runCreateDir(args []string) error {
	if len(args) != 2 {
		return errors.New("you should specify the directory you want to create and parent directory")
	}

	createDir := args[0]
	parentDir := args[1]

	content, err := ioutil.ReadFile("state.json")

	if err != nil {
		return errors.New("you are not authorized")
	}

	var credentialsMap map[string]string

	_ = json.Unmarshal(content, &credentialsMap)

	credentialsPath, ok := credentialsMap["credentialsPath"]
	if !ok {
		return errors.New("you are not authorized")
	}

	service, err := drive.Authorize(credentialsPath)

	if err != nil {
		return err
	}

	_, err = service.CreateDir(createDir, parentDir)

	return err
}
