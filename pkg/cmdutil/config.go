package cmdutil

import "github.com/bakurits/fshare-cli/pkg/cfg"

type Config struct {
	TokenPath     string
	Host          string
	MailStorePath string

	GoogleCredentialsPath string
	GoogleCredentials     cfg.GoogleCredentials
}
