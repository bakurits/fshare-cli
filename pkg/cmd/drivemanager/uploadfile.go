package drivemanager

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/drive"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type UploadFileCommand struct {
	AuthClient *auth.Client
}

// New: authorizeCmd represents the authorize command
func (c UploadFileCommand) New() *cobra.Command {
	var authorizeCmd = &cobra.Command{
		Use:   "uploadfile",
		Short: "accepts list of files and uploading in google drive",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runUploadFile(args)
		},
	}
	return authorizeCmd
}

func (c UploadFileCommand) runUploadFile(args []string) error {
	if len(args) == 0 {
		return errors.New("no file specified to upload")
	}

	service, err := drive.NewService(c.AuthClient.Client)
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
