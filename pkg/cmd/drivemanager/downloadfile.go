package drivemanager

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewDownloadCommand : downloadCmd represents the download command
func NewDownloadCommand() *cobra.Command {
	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "downloads file from google drive",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("too many arguments")
			}
			return download(args[0])
		},
	}
	return downloadCmd
}

// download : make download
func download(_ string) error {

	return nil
}
