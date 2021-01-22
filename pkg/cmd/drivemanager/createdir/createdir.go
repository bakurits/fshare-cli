package createdir

import (
	"github.com/bakurits/fshare-cli/pkg/drive"

	"github.com/bakurits/fshare-common/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CreateDirOptions options for createdir
type CreateDirOptions struct {
	name   string
	parent string
}

// CreateDirCommand stores dependencies for createdir command
type CreateDirCommand struct {
	AuthClient *auth.Client
}

// New : NewCreateDirCommand represents the creation of dir command
func (c CreateDirCommand) New() *cobra.Command {
	var opts CreateDirOptions

	// createdirCmd represents the createdir command
	var createdirCmd = &cobra.Command{
		Use:   "createdir",
		Short: "creation of directory in google drive",
		Long:  `creation of directory in google drive, first argument is dir you want to create, second argument is parent directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runCreateDir(opts)
		},
	}

	createdirCmd.Flags().StringVarP(&opts.name, "name", "n", "", "name of the directory")
	createdirCmd.Flags().StringVarP(&opts.parent, "parent", "p", "", "parent directory name")
	_ = createdirCmd.MarkFlagRequired("name")
	_ = createdirCmd.MarkFlagRequired("parent")

	return createdirCmd
}

func (c CreateDirCommand) runCreateDir(opts CreateDirOptions) error {
	if c.AuthClient == nil {
		return errors.New("unauthorized usser")
	}

	createDir := opts.name
	parentDir := opts.parent
	service, err := drive.NewService(c.AuthClient.Client)
	if err != nil {
		return err
	}

	_, err = service.CreateDir(createDir, parentDir)
	return err
}
