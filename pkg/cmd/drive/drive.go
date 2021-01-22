package drive

import (
	auth2 "github.com/bakurits/fshare-cli/pkg/cmd/drive/auth"
	"github.com/bakurits/fshare-cli/pkg/cmd/drive/createdir"
	"github.com/bakurits/fshare-cli/pkg/cmd/drive/download"
	"github.com/bakurits/fshare-cli/pkg/cmd/drive/list"
	"github.com/bakurits/fshare-cli/pkg/cmd/drive/uploadfile"
	"github.com/bakurits/fshare-cli/pkg/cmdutil"
	"github.com/bakurits/fshare-common/auth"
	"github.com/spf13/cobra"
)

func New(conf *cmdutil.Config, authClient *auth.Client) *cobra.Command {
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
