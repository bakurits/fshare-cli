package drivemanager

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"

	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/spf13/cobra"
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

	for _, args := range args {
		err = uploadFile(args, service)
		if err != nil {
			return err
		}
	}

	return nil
}

func uploadFile(filepath string, s *drive.Service) error {
	f, err := os.Open(filepath)

	if err != nil {
		return errors.New("file does not exist " + filepath)
	}

	tp := getFilePathMimeType(filepath)

	_, err = s.CreateFile(filepath, tp, f, "root")

	defer f.Close()

	return err
}

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}
