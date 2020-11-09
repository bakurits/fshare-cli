package drive

import "github.com/bakurits/fileshare/pkg/cfg"

func getTestConfig() (cfg.Config, error) {
	conf := cfg.Config{
		GoogleCredentialsPath: "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/web_app_credentials.json",
		TokenPath:             "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/token.json",
	}
	err := cfg.LoadGoogleCredentials(&conf)
	if err != nil {
		return cfg.Config{}, err
	}
	return conf, nil
}
