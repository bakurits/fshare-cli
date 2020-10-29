package drivemanager

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/bakurits/fileshare/pkg/testutils"
	"github.com/spf13/cobra"
)

// NewListCommand : generates of command list command
func NewListCommand() *cobra.Command {

	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "show list of files",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runList()
			return err
		},
	}

	return listCmd
}

func runList() error {

	service, err := drive.Authorize(testutils.RootDir() + "/credentials")

	if err != nil {
		return err
	}

	list := service.List()

	for i := range list {
		fmt.Println(list[i].Name)
	}

	return nil
}
