package mail

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/pkg/errors"
	"os"
)

func getAuthClient() (*auth.Client, error) {
	appCfg, err := getConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to find app config")
	}
	authCfg := auth.GetConfig(appCfg)
	client, err := authCfg.ClientFromTokenFile(appCfg.TokenPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create auth client")
	}
	return client, nil
}

func getConfig() (cfg.Config, error) {
	cfgFile := os.Getenv("config")
	if cfgFile == "" {
		cfgFile = "config.json"
	}
	return cfg.GetConfig(cfgFile)
}
