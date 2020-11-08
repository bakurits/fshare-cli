package drivemanager

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/bakurits/fileshare/pkg/testutils"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	number int
}

// NewListCommand : generates of command list command
func NewListCommand() *cobra.Command {
	var opts ListOptions

	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "show list of files",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runList(opts)
			return err
		},
	}
	listCmd.Flags().IntVarP(&opts.number, "number", "n", 20, "number of files")

	return listCmd
}

func runList(opts ListOptions) error {
	service, err := drive.Authorize(testutils.RootDir() + "/credentials")
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
