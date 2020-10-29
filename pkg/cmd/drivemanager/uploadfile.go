package drivemanager

import (
	"errors"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/bakurits/fileshare/pkg/testutils"
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

	service, err := drive.Authorize(testutils.RootDir() + "/credentials")

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
