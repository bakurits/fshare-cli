package mail

import (
	"path/filepath"

	"github.com/bakurits/fshare-cli/pkg/auth"
	"github.com/bakurits/fshare-cli/pkg/gmail"

	"github.com/spf13/cobra"
)

type SendMAilOptions struct {
	FromMail       string
	ToMail         string
	AttachmentPath string
	Subject        string
	Content        string
}

type SendAttachmentCommand struct {
	AuthClient *auth.Client
}

// New: generates of command createdir
func (c SendAttachmentCommand) New() *cobra.Command {
	var opts SendMAilOptions

	// sendmailCmd represents the sendmail command
	var sendmailCmd = &cobra.Command{
		Use:   "sendattachment",
		Short: "sending gmail or attachment to users",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runSendMail(opts)
		},
	}

	sendmailCmd.Flags().StringVar(&opts.FromMail, "from", "", "sender gmail")
	sendmailCmd.Flags().StringVar(&opts.ToMail, "to", "", "receiver gmail")
	sendmailCmd.Flags().StringVarP(&opts.AttachmentPath, "path", "p", "", "path to attachment")
	sendmailCmd.Flags().StringVar(&opts.Content, "content", "", "content message, default empty text")
	sendmailCmd.Flags().StringVar(&opts.Subject, "subject", "", "Subject gmail, default empty text")

	_ = sendmailCmd.MarkFlagRequired("from")
	_ = sendmailCmd.MarkFlagRequired("to")
	_ = sendmailCmd.MarkFlagRequired("path")

	return sendmailCmd
}

// runSendMail : sending gmail command
func (c SendAttachmentCommand) runSendMail(opts SendMAilOptions) error {
	fileDir, fileName := filepath.Split(opts.AttachmentPath)
	srv, err := gmail.NewService(c.AuthClient.Client)
	if err != nil {
		return err
	}
	messageWithAttachment := gmail.CreateMessageWithAttachment("me", opts.ToMail, opts.Subject, opts.Content, fileDir, fileName)
	err = srv.SendMessage("me", messageWithAttachment)
	return err
}
