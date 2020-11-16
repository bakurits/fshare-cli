package drivemanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/webapp/server"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type AuthorizeOptions struct {
	email string
}

// NewAuthorizeCommand : authorizeCmd represents the authorize command
func NewAuthorizeCommand() *cobra.Command {
	var opts AuthorizeOptions
	var authorizeCmd = &cobra.Command{
		Use:   "auth",
		Short: "make authorization in google drive with credentials with given directory which holds a file credentialsMail.json",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return authorize(opts)
		},
	}

	authorizeCmd.Flags().StringVarP(&opts.email, "email", "m", "", "email")
	authorizeCmd.MarkFlagRequired("email")

	return authorizeCmd
}

// storeToken : make a get request to a server for getting token
func storeToken(email string, password string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, cfg.Host+server.GetTokenEndpoint, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(email, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var tok oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return err
	}

	if err := auth.SaveToken(cfg.TokenPath, &tok); err != nil {
		return err
	}
	return nil
}

// authorize : make authorization
func authorize(opts AuthorizeOptions) error {
	prompt := fmt.Sprintf("Enter password:\n")
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	return storeToken(opts.email, password)
}
