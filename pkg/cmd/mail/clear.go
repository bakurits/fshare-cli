package mail

import (
	"github.com/bakurits/fshare-cli/pkg/mailstore"

	"github.com/spf13/cobra"
)

// SendAttachmentCommand stores dependencies for send attachment command
type ClearMailStoreCommand struct {
	MailStorePath string
}

func (c ClearMailStoreCommand) New() *cobra.Command {
	// sendmailCmd represents the sendmail command
	var clearMailCmd = &cobra.Command{
		Use:   "clear",
		Short: "clear mail store",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runClearMail()
		},
	}

	return clearMailCmd
}

func (c ClearMailStoreCommand) runClearMail() error {
	return mailstore.DeleteMails(c.MailStorePath)
}
