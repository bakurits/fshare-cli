package mail

import (
	"github.com/bakurits/fileshare/pkg/gmail"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"path/filepath"
)

type SendMAilOptions struct {
	FromMail       string
	ToMail         string
	AttachmentPath string
	Subject        string
	Content        string
}

// NewCreateDirCommand : generates of command createdir
func NewSendAttachmentCommand() *cobra.Command {
	var opts SendMAilOptions

	// sendmailCmd represents the sendmail command
	var sendmailCmd = &cobra.Command{
		Use:   "sendattachment",
		Short: "sending gmail or attachment to users",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSendMail(opts)
		},
	}

	sendmailCmd.Flags().StringVar(&opts.FromMail, "from", "", "sender gmail")
	sendmailCmd.Flags().StringVar(&opts.ToMail, "to", "", "receiver gmail")
	sendmailCmd.Flags().StringVarP(&opts.AttachmentPath, "path", "p", "", "path to attachment")
	sendmailCmd.Flags().StringVar(&opts.Content, "content", "", "content message, default empty text")
	sendmailCmd.Flags().StringVar(&opts.Subject, "subject", "", "Subject gmail, default empty text")

	sendmailCmd.MarkFlagRequired("from")
	sendmailCmd.MarkFlagRequired("to")
	sendmailCmd.MarkFlagRequired("path")

	return sendmailCmd
}

// runSendMail : sending gmail command
func runSendMail(opts SendMAilOptions) error {
	fileDir, fileName := filepath.Split(opts.AttachmentPath)
	authClient, err := getAuthClient()
	if err != nil {
		return errors.Wrap(err, "auth error")
	}
	srv, err := gmail.NewService(authClient.Client)
	if err != nil {
		return err
	}
	messageWithAttachment := gmail.CreateMessageWithAttachment("me", opts.ToMail, opts.Subject, opts.Content, fileDir, fileName)
	err = srv.SendMessage("me", messageWithAttachment)
	return err
}
