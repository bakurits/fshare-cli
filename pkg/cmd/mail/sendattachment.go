package mail

import (
	"path/filepath"

	"github.com/bakurits/fshare-cli/pkg/gmail"
	"github.com/bakurits/fshare-cli/pkg/mailstore"
	"github.com/bakurits/fshare-common/auth"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// SendMailOptions options for send attachment command
type SendMailOptions struct {
	FromMail       string
	ToMail         string
	AttachmentPath string
	Subject        string
	Content        string
}

// SendAttachmentCommand stores dependencies for send attachment command
type SendAttachmentCommand struct {
	AuthClient    *auth.Client
	MailStorePath string
}

// New generates of command createdir
func (c SendAttachmentCommand) New() *cobra.Command {
	var opts SendMailOptions

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
	//_ = sendmailCmd.MarkFlagRequired("to")
	_ = sendmailCmd.MarkFlagRequired("path")

	return sendmailCmd
}

func containMail(mails []string, mail string) bool {
	for _, v := range mails {
		if v == mail {
			return true
		}
	}

	return false
}

func (c SendAttachmentCommand) saveMail(toMail string) error {
	mails, err := mailstore.ReadMails(c.MailStorePath)
	if err != nil {
		return err
	}

	if containMail(mails, toMail) {
		return nil
	}

	return mailstore.WriteMail(toMail, c.MailStorePath)
}

func (c SendAttachmentCommand) chooseToMAil() (string, error) {
	mails, err := mailstore.ReadMails(c.MailStorePath)
	if err != nil {
		return "", err
	}

	if len(mails) == 0 {
		return "", nil
	}

	var qs = []*survey.Question{
		{
			Name: "mail",
			Prompt: &survey.Select{
				Message: "Choose a mail:",
				Options: mails,
				Default: mails[0],
			},
		},
	}

	answers := struct {
		Mail string
	}{}

	err = survey.Ask(qs, &answers)
	if err != nil {
		return "", err
	}

	return answers.Mail, nil
}

// runSendMail : sending gmail command
func (c SendAttachmentCommand) runSendMail(opts SendMailOptions) error {
	fileDir, fileName := filepath.Split(opts.AttachmentPath)
	srv, err := gmail.NewService(c.AuthClient.Client)
	if err != nil {
		return err
	}

	if opts.ToMail != "" {
		err = c.saveMail(opts.ToMail)
		if err != nil {
			return err
		}
	} else {
		opts.ToMail, err = c.chooseToMAil()
		if err != nil {
			return err
		}
	}

	messageWithAttachment := gmail.CreateMessageWithAttachment("me", opts.ToMail, opts.Subject, opts.Content, fileDir, fileName)
	err = srv.SendMessage("me", messageWithAttachment)
	return err
}
