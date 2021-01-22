package root

import (
	"github.com/bakurits/fshare-cli/pkg/cfg"
	"github.com/bakurits/fshare-cli/pkg/cmd/drivemanager"
	drivemanager2 "github.com/bakurits/fshare-cli/pkg/cmd/mail"
	"github.com/bakurits/fshare-cli/pkg/cmdutil"
	"github.com/bakurits/fshare-common/auth"
	"github.com/spf13/cobra"
)

func NewCmdRoot(conf *cmdutil.Config, authClient *auth.Client) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   cfg.AppName,
		Short: "",
		Long:  ``,
	}
	initCommands(rootCmd, conf, authClient)
	return rootCmd
}

func initCommands(rootCmd *cobra.Command, conf *cmdutil.Config, authClient *auth.Client) {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(drivemanager.NewDrive(conf, authClient))
	rootCmd.AddCommand(drivemanager2.NewMail(conf, authClient))
	rootCmd.AddCommand(NewCmdCompletion())
}
