package drivemanager

import (
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
func getToken(email string, password string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/token", nil)
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
	return string(body), nil
}

func writeToken(body string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}
	f, err := os.Create(cfg.CredentialsDir + "/token.json")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	return err
}

// authorize : make authorization
func authorize(opts AuthorizeOptions) error {
	prompt := fmt.Sprintf("Enter password:\n")
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	body, err := getToken(opts.email, password)
	if err != nil {
		return err
	}
	err = writeToken(body)
	return err
}
