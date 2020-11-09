package drivemanager

import (
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCreateDirCommand : NewCreateDirCommand represents the creation of dir command
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

	authClient, err := getAuthClient()
	if err != nil {
		return errors.Wrap(err, "auth error")
	}

	service, err := drive.NewService(authClient.Client)
	if err != nil {
		return err
	}

	_, err = service.CreateDir(createDir, parentDir)

	return err
}
