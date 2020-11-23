package drive

import (
	"github.com/bakurits/fshare-cli/pkg/auth"
	"github.com/bakurits/fshare-cli/pkg/cfg"
)

func getTestClient() (*auth.Client, error) {
	tokenPath := "testcredentials/token.json"
	credentialsPath := "testcredentials/google_credentials.json"
	cred, err := cfg.LoadGoogleCredentials(credentialsPath, cfg.CredentialTypeWeb)
	if err != nil {
		return nil, err
	}

	authClient, err := auth.
		GetConfig(cred.ClientID, cred.ClientSecret, "http://localhost").
		ClientFromTokenFile(tokenPath)
	if err != nil {
		return nil, err
	}
	return authClient, nil
}
