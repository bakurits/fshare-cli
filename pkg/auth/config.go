package auth

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

type GoogleCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	ProjectID string `json:"project_id"`
}

type WebConfig struct {
	StaticFileDir  string
	CredentialsDir string

	ConnectionString string
	DBDialect        string

	Server string
	Port   string

	GoogleCredentialsPath string
	GoogleCredentials     GoogleCredentials
}

type CmdConfig struct {
	CredentialsDir string
	TokenPath      string
	Host           string

	GoogleCredentialsPath string
	GoogleCredentials     GoogleCredentials
}

func LoadGoogleCredentials(GoogleCredentialsPath string) (GoogleCredentials, error) {
	b, err := ioutil.ReadFile(GoogleCredentialsPath)
	if err != nil {
		return GoogleCredentials{}, errors.Wrap(err, "unable to read client secret file")
	}
	jsn := struct {
		Web GoogleCredentials `json:"web"`
	}{}
	err = json.Unmarshal(b, &jsn)
	if err != nil {
		return GoogleCredentials{}, errors.Wrap(err, "unable to parse google credentials")
	}
	return jsn.Web, nil
}
