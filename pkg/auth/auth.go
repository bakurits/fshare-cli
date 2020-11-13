package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bakurits/fileshare/pkg/cfg"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
)

// Scopes needed for API
var Scopes = []string{
	drive.DriveScope,
	gmail.GmailMetadataScope,
	gmail.GmailSendScope,
}

type Config struct {
	appConfig  cfg.Config
	authConfig oauth2.Config
}

func GetConfig(conf cfg.Config) *Config {
	return &Config{
		appConfig: conf,
		authConfig: oauth2.Config{
			ClientID:     conf.GoogleCredentials.ClientID,
			ClientSecret: conf.GoogleCredentials.ClientSecret,
			RedirectURL:  conf.Server + conf.Port + "/auth",
			Scopes:       Scopes,
			Endpoint:     google.Endpoint,
		},
	}
}

func (cfg *Config) AuthCodeURL(state string) string {
	return cfg.authConfig.AuthCodeURL(state)
}

type Client struct {
	*http.Client
	Token *oauth2.Token

	Email string
}

// ClientFromCode returns auth client from code
func (cfg *Config) ClientFromCode(code string) (*Client, error) {
	tok, err := cfg.authConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.Wrap(err, "error while generating token")
	}
	_ = saveToken("newToken.json", tok)

	return cfg.ClientFromToken(tok)
}

// ClientFromTokenFile return auth token from token file
func (cfg *Config) ClientFromTokenFile(path string) (*Client, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file")
	}
	v := oauth2.Token{}
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal token")
	}
	return cfg.ClientFromToken(&v)
}

// ClientFromToken return auth client from token
func (cfg *Config) ClientFromToken(tok *oauth2.Token) (*Client, error) {
	client := cfg.authConfig.Client(context.Background(), tok)
	resp, err := client.Get("https://gmail.googleapis.com/gmail/v1/users/me/profile")
	if err != nil {
		return nil, errors.Wrap(err, "error while getting user information")
	}
	defer func() { _ = resp.Body.Close() }()

	user := struct {
		Email string `json:"emailAddress"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, errors.Wrap(err, "error while retrieving user email")
	}

	return &Client{
		Client: client,
		Token:  tok,

		Email: user.Email,
	}, nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrap(err, "unable to cache oauth token")
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err = json.NewEncoder(f).Encode(token); err != nil {
		return errors.Wrap(err, "unable to encode json data")
	}
	return nil
}
