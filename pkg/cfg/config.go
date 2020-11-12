package cfg

import (
	"encoding/json"
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Config struct {
	StaticFileDir string

	ConnectionString string
	DBDialect        string

	Server string
	Port   string

	GoogleCredentialsPath string
	GoogleCredentials     GoogleCredentials

	TokenPath string
}

type GoogleCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	ProjectID string `json:"project_id"`
}

func GetConfig(path string) (Config, error) {
	var conf Config
	err := configor.Load(&conf, path)
	if err != nil {
		return Config{}, err
	}

	err = LoadGoogleCredentials(&conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func LoadGoogleCredentials(config *Config) error {
	b, err := ioutil.ReadFile(config.GoogleCredentialsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read client secret file")
	}
	jsn := struct {
		Web GoogleCredentials `json:"web"`
	}{}
	err = json.Unmarshal(b, &jsn)
	if err != nil {
		return errors.Wrap(err, "unable to parse google credentials")
	}
	config.GoogleCredentials = jsn.Web
	return nil
}
