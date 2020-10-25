package cmd

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/cmd/customoperations"
	"github.com/bakurits/fileshare/pkg/cmd/drivemanager"
	"github.com/bakurits/fileshare/pkg/cmd/mail"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

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

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fileshare.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(customoperations.NewAddCommand())
	rootCmd.AddCommand(drivemanager.NewUploadFileCommand())
	rootCmd.AddCommand(drivemanager.NewCreateDirCommand())
	rootCmd.AddCommand(drivemanager.NewAuthorizeCommand())
	rootCmd.AddCommand(drivemanager.NewListCommand())
	rootCmd.AddCommand(drivemanager.NewDownloadCommand())
	rootCmd.AddCommand(mail.NewSendAttachmentCommand())
}

// initConfig : initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".fileshare" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".fileshare")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
