package drivemanager

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/bakurits/fileshare/pkg/testutils"
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
			if len(args) < 1 {
				return errors.New("file name must be specified")
			}
			if len(args) > 1 {
				return errors.New("too many arguments")
			}
			return download(args[0])
		},
	}
	return downloadCmd
}

// download : make download
func download(name string) error {
	client, err := auth.GetHTTPClient(testutils.RootDir() + "/credentials")

	srv, err := drive.NewService(client)
	if err != nil {
		return errors.Wrap(err, "unexpected error")
	}

	f, err := srv.Get(name)
	if err != nil {
		return errors.Wrap(err, "fetching error")
	}

	err = srv.Download(f)
	if err != nil {
		return errors.Wrap(err, "downloading error")
	}
	return nil
}
