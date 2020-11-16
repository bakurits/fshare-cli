package drivemanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"
)

type AuthorizeOptions struct {
	email string
}

// NewAuthorizeCommand : authorizeCmd represents the authorize command
func NewAuthorizeCommand() *cobra.Command {
	var opts AuthorizeOptions
	var authorizeCmd = &cobra.Command{
		Use:   "authorize",
		Short: "make authorization in google drivemanager with credentials with given directory which holds a file credentialsMail.json",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return authorize(opts)
		},
	}

	authorizeCmd.Flags().StringVarP(&opts.email, "email", "m", "", "email")
	authorizeCmd.MarkFlagRequired("email")

	return authorizeCmd
}

// getToken : make a get request to a server for getting token
func getToken(email string, password string, getUrl string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(email, password)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}
	return string(body), nil
}

func writeToken(body string, CredentialsDir string) error {
	f, err := os.Create(CredentialsDir + "/token.json")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	return err
}

// authorize : make authorization
func authorize(opts AuthorizeOptions) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}
	prompt := fmt.Sprintf("Enter password:\n")
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	body, err := getToken(opts.email, password, cfg.GetUrl)
	if err != nil {
		return err
	}
	err = writeToken(body, cfg.CredentialsDir)
	return err
}
