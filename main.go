package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bakurits/fshare-cli/pkg/cfg"
	"github.com/bakurits/fshare-cli/pkg/root"

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

func readConfig() *root.Config {
	var conf root.Config
	err := cfg.GetConfig(&conf)
	if err != nil {
		log.Fatal("Can't find config")
	}
	cred, err := cfg.LoadGoogleCredentials(conf.GoogleCredentialsPath, cfg.CredentialTypeDesktop)
	if err != nil {
		log.Fatal("Unable to parse credentials")
	}
	conf.GoogleCredentials = cred
	return &conf
}

func getAuthClient(conf *root.Config) *auth.Client {
	var err error
	authClient, err := auth.
		GetConfig(conf.GoogleCredentials.ClientID, conf.GoogleCredentials.ClientSecret, "http://localhost").
		ClientFromTokenFile(conf.TokenPath)
	if err != nil {
		return nil
	}
	return authClient
}

func initConfig() (*root.Config, *auth.Client) {
	conf := readConfig()
	authClient := getAuthClient(conf)
	return conf, authClient
}

func main() {
	conf, authClient := initConfig()
	rootCmd := root.NewCmdRoot(conf, authClient)
	Execute(rootCmd)
}
