package drivemanager

import (
	"github.com/bakurits/fshare-cli/pkg/drive"

	"github.com/bakurits/fshare-common/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// UploadFileCommand stores dependencies for upload file command
type UploadFileCommand struct {
	AuthClient *auth.Client
}

// New authorizeCmd represents the authorize command
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

	if c.AuthClient == nil {
		return errors.New("Unauthorized User")
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
