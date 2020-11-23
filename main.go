package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bakurits/fshare-cli/pkg/cfg"
	"github.com/bakurits/fshare-cli/pkg/cmd"

	"github.com/bakurits/fshare-common/auth"
	"github.com/spf13/cobra"
)

// Execute executes cmd command
func Execute(rootCmd *cobra.Command) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var authClient *auth.Client

func readConfig() *cmd.Config {
	var conf cmd.Config
	err := cfg.GetConfig(&conf)
	if err != nil {
		log.Fatal(err)
	}
	cred, err := cfg.LoadGoogleCredentials(conf.GoogleCredentialsPath, cfg.CredentialTypeDesktop)
	if err != nil {
		log.Fatal(err)
	}
	conf.GoogleCredentials = cred
	return &conf
}

func getAuthClient(conf *cmd.Config) *auth.Client {
	var err error
	authClient, err = auth.
		GetConfig(conf.GoogleCredentials.ClientID, conf.GoogleCredentials.ClientSecret, "http://localhost").
		ClientFromTokenFile(conf.TokenPath)
	if err != nil {
		return nil
	}
	return authClient
}

// initConfig : initConfig reads in config file and ENV variables if set.
func initConfig() (*cmd.Config, *auth.Client) {
	conf := readConfig()
	authClient := getAuthClient(conf)
	return conf, authClient
}

func main() {
	conf, authClient := initConfig()
	rootCmd := cmd.NewCmdRoot(conf, authClient)
	Execute(rootCmd)
}
