package drive

import "github.com/bakurits/fileshare/pkg/auth"

func getTestConfig() (auth.CmdConfig, error) {
	conf := auth.CmdConfig{
		GoogleCredentialsPath: "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/web_app_credentials.json",
		TokenPath:             "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/token.json",
	}
	web, err := auth.LoadGoogleCredentials(conf.GoogleCredentialsPath)
	conf.GoogleCredentials = web
	if err != nil {
		return auth.CmdConfig{}, err
	}
	return conf, nil
}
