package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
)

// GetHTTPClient return http client authorized in google
func GetHTTPClient(credentialsDir string) (*http.Client, error) {
	b, err := ioutil.ReadFile(credentialsDir + "/credentials.json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to read client secret file")
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope, gmail.GmailMetadataScope)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse client secret file to config")
	}
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := credentialsDir + "/token.json"
	tok, err := tokenFromFile(tokFile)
	if err == nil {
		return config.Client(context.Background(), tok), nil
	}

	tok, err = getTokenFromWeb(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get token from web")
	}
	if err = saveToken(tokFile, tok); err != nil {
		return nil, errors.Wrap(err, "unable to save token")
	}

	return config.Client(context.Background(), tok), nil

}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, errors.Wrap(err, "unable to read authorization code")
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve token from web ")
	}
	return tok, nil
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
