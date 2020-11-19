package cfg

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
)

const (
	AppName = "fileshare"
)

type GoogleCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	ProjectID string `json:"project_id"`
}

func GetConfig(conf interface{}) error {
	viper.SetConfigName("config")                          // name of config file (without extension)
	viper.SetConfigType("json")                            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", AppName))  // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", AppName)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                               // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {           // Handle errors reading the config file
		return errors.Wrap(err, "error while reading config file")
	}

	if err := viper.Unmarshal(conf); err != nil {
		return errors.Wrap(err, "error while reading config file")
	}
	return nil
}

type CredentialType bool

const (
	CredentialTypeWeb     CredentialType = true
	CredentialTypeDesktop CredentialType = false
)

func LoadGoogleCredentials(credentialsPath string, tp CredentialType) (GoogleCredentials, error) {
	b, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		return GoogleCredentials{}, errors.Wrap(err, "unable to read client secret file")
	}

	switch tp {
	case CredentialTypeWeb:
		jsn := struct {
			Web GoogleCredentials `json:"web"`
		}{}
		err = json.Unmarshal(b, &jsn)
		if err != nil {
			return GoogleCredentials{}, errors.Wrap(err, "unable to parse google credentials")
		}
		return jsn.Web, nil
	case CredentialTypeDesktop:
		jsn := struct {
			Installed GoogleCredentials `json:"installed"`
		}{}
		err = json.Unmarshal(b, &jsn)
		if err != nil {
			return GoogleCredentials{}, errors.Wrap(err, "unable to parse google credentials")
		}
		return jsn.Installed, nil
	}
	return GoogleCredentials{}, errors.New("unable to parse credentials")
}
