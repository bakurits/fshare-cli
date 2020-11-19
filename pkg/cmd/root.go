package cmd

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/bakurits/fileshare/pkg/cmd/drivemanager"
	"github.com/bakurits/fileshare/pkg/cmd/mail"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfgFile string

type Config struct {
	TokenPath string
	Host      string

	GoogleCredentialsPath string
	GoogleCredentials     cfg.GoogleCredentials
}

var conf Config
var authClient *auth.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fileshare",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	initConfig()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(drivemanager.AuthorizeCommand{Host: conf.Host, TokenPath: conf.TokenPath}.New())
	rootCmd.AddCommand(drivemanager.UploadFileCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.CreateDirCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.ListCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(drivemanager.DownloadCommand{AuthClient: authClient}.New())
	rootCmd.AddCommand(mail.SendAttachmentCommand{AuthClient: authClient}.New())

}

func readConfig() {
	err := cfg.GetConfig(&conf)
	if err != nil {
		log.Fatal(err)
	}
	cred, err := cfg.LoadGoogleCredentials(conf.GoogleCredentialsPath, cfg.CredentialTypeDesktop)
	if err != nil {
		log.Fatal(err)
	}
	conf.GoogleCredentials = cred
}

func getAuthClient() error {
	var err error
	authClient, err = auth.
		GetConfig(conf.GoogleCredentials.ClientID, conf.GoogleCredentials.ClientSecret, "http://localhost").
		ClientFromTokenFile(conf.TokenPath)
	if err != nil {
		return err
	}
	return nil
}

// initConfig : initConfig reads in config file and ENV variables if set.
func initConfig() {
	readConfig()
	_ = getAuthClient()
}
