package mail

import (
	"github.com/bakurits/fshare-cli/pkg/cmdutil"
	"github.com/bakurits/fshare-common/auth"

	"github.com/spf13/cobra"
)

func New(conf *cmdutil.Config, authClient *auth.Client) *cobra.Command {
	mailCmd := &cobra.Command{
		Use:   "mail <command>",
		Short: "send",
		Long:  ``,
	}

	mailCmd.AddCommand(SendAttachmentCommand{AuthClient: authClient, MailStorePath: conf.MailStorePath}.New())
	mailCmd.AddCommand(ClearMailStoreCommand{MailStorePath: conf.MailStorePath}.New())
	return mailCmd
}
