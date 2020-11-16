package drivemanager

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/pkg/errors"
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
	cfgFile := "C:\\Users\\Giorgi\\GolandProjects\\fileshare\\cmd\\webapp\\config.json"
	if cfgFile == "" {
		cfgFile = "config.json"
	}
	return cfg.GetConfig(cfgFile)
}
