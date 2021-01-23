package drive

import (
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

	driveCmd.AddCommand(AuthorizeCommand{Host: conf.Host, TokenPath: conf.TokenPath}.New())
	driveCmd.AddCommand(UploadFileCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(CreateDirCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(ListCommand{AuthClient: authClient}.New())
	driveCmd.AddCommand(DownloadCommand{AuthClient: authClient}.New())

	return driveCmd
}
