package drive

import (
	"fmt"

	"github.com/bakurits/fshare-cli/pkg/drive"

	"github.com/bakurits/fshare-common/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ListOptions options for list command
type ListOptions struct {
	number int
}

// ListCommand stores dependencies for list command
type ListCommand struct {
	AuthClient *auth.Client
}

// New : generates of command list command
func (c ListCommand) New() *cobra.Command {
	var opts ListOptions

	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "show list of files",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.runList(opts)
			return err
		},
	}
	listCmd.Flags().IntVarP(&opts.number, "number", "n", 20, "number of files")

	return listCmd
}

func (c ListCommand) runList(opts ListOptions) error {
	if c.AuthClient == nil {
		return errors.New("unauthorized user")
	}

	service, err := drive.NewService(c.AuthClient.Client)
	if err != nil {
		return err
	}

	list, err := service.List(opts.number)
	if err != nil {
		return err
	}
	for i := range list {
		fmt.Println(list[i].Name)
	}
	return nil
}
