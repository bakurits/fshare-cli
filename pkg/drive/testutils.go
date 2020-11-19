package drive

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
)

func getTestClient() (*auth.Client, error) {
	tokenPath := "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/token.json"
	credentialsPath := "/home/bakurits/Documents/Programming/Utilities/fileshare/credentials/google_credentials.json"
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
