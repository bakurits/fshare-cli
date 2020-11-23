package cmd

import (
	"github.com/bakurits/fshare-cli/pkg/cfg"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager"
	"github.com/bakurits/fshare-cli/pkg/cmd/mail"
	"github.com/bakurits/fshare-common/auth"

	"github.com/spf13/cobra"
)

type Config struct {
	TokenPath string
	Host      string

	GoogleCredentialsPath string
	GoogleCredentials     cfg.GoogleCredentials
}

func NewCmdRoot(conf *Config, authClient *auth.Client) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "fileshare",
		Short: "",
		Long:  ``,
	}
	initCommands(rootCmd, conf, authClient)
	return rootCmd
}

func initCommands(rootCmd *cobra.Command, conf *Config, authClient *auth.Client) {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(drivemanager.AuthorizeCommand{Host: conf.Host, TokenPath: conf.TokenPath}.New())
	rootCmd.AddCommand(drivemanager.UploadFileCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.CreateDirCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.ListCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.DownloadCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(mail.SendAttachmentCommand{AuthClient: authClient}.New())
}
