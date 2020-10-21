package drivemanager

import (
	"encoding/json"
	"errors"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// NewUploadFileCommand : authorizeCmd represents the authorize command
func NewUploadFileCommand() *cobra.Command {
	var authorizeCmd = &cobra.Command{
		Use:   "uploadfile",
		Short: "accepts list of files and uploading in google drivemanager",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUploadFile(args)
		},
	}
	return authorizeCmd
}

func runUploadFile(args []string) error {

	if len(args) == 0 {
		return errors.New("no file specified to upload")
	}
	// Read the file
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

	for _, arg := range args {
		_, err = service.Upload(arg)
		if err != nil {
			return err
		}
	}

	return nil
}
