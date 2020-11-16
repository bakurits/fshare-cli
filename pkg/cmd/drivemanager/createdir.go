package drivemanager

import (
	"github.com/bakurits/fileshare/pkg/drive"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type CreateDirOptions struct {
	name   string
	parent string
}

// NewCreateDirCommand : NewCreateDirCommand represents the creation of dir command
func NewCreateDirCommand() *cobra.Command {
	var opts CreateDirOptions

	// createdirCmd represents the createdir command
	var createdirCmd = &cobra.Command{
		Use:   "createdir",
		Short: "creation of directory in google drive",
		Long:  `creation of directory in google drive, first argument is dir you want to create, second argument is parent directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateDir(opts)
		},
	}

	createdirCmd.Flags().StringVarP(&opts.name, "name", "n", "", "name of the directory")
	createdirCmd.Flags().StringVarP(&opts.parent, "parent", "p", "", "parent directory name")
	createdirCmd.MarkFlagRequired("name")
	createdirCmd.MarkFlagRequired("parent")

	return createdirCmd
}

func runCreateDir(opts CreateDirOptions) error {
	createDir := opts.name
	parentDir := opts.parent
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
