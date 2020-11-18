package drivemanager

import (
	"log"
	"os"

	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"

	"github.com/pkg/errors"
)

func getAuthClient() (*auth.Client, error) {
	appCfg, err := getConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to find app config")
	}
	authCfg := auth.GetCmdConfig(appCfg)
	client, err := authCfg.ClientFromTokenFile(appCfg.TokenPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create auth client")
	}
	return client, nil
}

func getConfig() (auth.CmdConfig, error) {
	cfgFile := os.Getenv("config")
	if cfgFile == "" {
		cfgFile = "config.json"
	}

	var cmdConfig auth.CmdConfig
	err := cfg.GetConfig(cfgFile, &cmdConfig)
	if err != nil {
		return auth.CmdConfig{}, err
	}

	web, err := auth.LoadGoogleCredentials(cmdConfig.GoogleCredentialsPath)
	if err != nil {
		log.Fatal(err)
	}
	cmdConfig.GoogleCredentials = web

	return cmdConfig, nil
}
