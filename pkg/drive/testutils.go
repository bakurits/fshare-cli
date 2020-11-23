package drive

import (
	"context"
	"time"

	"github.com/bakurits/fshare-common/auth"
	"github.com/sethvargo/go-envconfig"
	"golang.org/x/oauth2"
)

type config struct {
	ClientID     string `env:"TEST_CLIENT_ID"`
	ClientSecret string `env:"TEST_CLIENT_SECRET"`
	ProjectID    string `env:"TEST_PROJECT_ID"`
	AccessToken  string `env:"TEST_ACCESS_TOKEN"`
	TokenType    string `env:"TEST_TOKEN_TYPE"`
	RefreshToken string `env:"TEST_REFRESH_TOKEN"`
	Expiry       string `env:"TEST_EXPIRY"`
}

func getTestClient() (*auth.Client, error) {
	var conf config
	if err := envconfig.Process(context.Background(), &conf); err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, conf.Expiry)
	if err != nil {
		return nil, err
	}

	authClient, err := auth.
		GetConfig(conf.ClientID, conf.ClientSecret, "http://localhost").
		ClientFromToken(&oauth2.Token{
			AccessToken:  conf.AccessToken,
			TokenType:    conf.TokenType,
			RefreshToken: conf.RefreshToken,
			Expiry:       t,
		})

	if err != nil {
		return nil, err
	}
	return authClient, nil
}
