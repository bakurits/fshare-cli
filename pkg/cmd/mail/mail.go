package drivemanager

import (
	"github.com/bakurits/fshare-cli/pkg/cmd/mail/clear"
	"github.com/bakurits/fshare-cli/pkg/cmd/mail/sendmail"
	"github.com/bakurits/fshare-cli/pkg/cmdutil"
	"github.com/bakurits/fshare-common/auth"
	"github.com/spf13/cobra"
)

func NewMail(conf *cmdutil.Config, authClient *auth.Client) *cobra.Command {
	mailCmd := &cobra.Command{
		Use:   "mail <command>",
		Short: "send",
		Long:  ``,
	}

	mailCmd.AddCommand(sendmail.SendAttachmentCommand{AuthClient: authClient, MailStorePath: conf.MailStorePath}.New())
	mailCmd.AddCommand(clear.ClearMailStoreCommand{MailStorePath: conf.MailStorePath}.New())
	return mailCmd
}
