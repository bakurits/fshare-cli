package drivemanager

import (
	auth2 "github.com/bakurits/fshare-cli/pkg/cmd/drivemanager/auth"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager/createdir"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager/download"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager/list"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager/uploadfile"
	"github.com/bakurits/fshare-cli/pkg/cmdutil"
	"github.com/bakurits/fshare-common/auth"
	"github.com/spf13/cobra"
)

func NewDrive(conf *cmdutil.Config, authClient *auth.Client) *cobra.Command {
	driveCmd := &cobra.Command{
		Use:   "drive <command>",
		Short: "download, createdir, uploadfile, list, auth",
		Long:  ``,
	}

	driveCmd.AddCommand(auth2.AuthorizeCommand{Host: conf.Host, TokenPath: conf.TokenPath}.New())
	driveCmd.AddCommand(uploadfile.UploadFileCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(createdir.CreateDirCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(list.ListCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(download.DownloadCommand{AuthClient: authClient}.New())

	return driveCmd
}
