package drivemanager

import (
	"encoding/json"
	"net/http"

	"github.com/bakurits/fshare-cli/pkg/cfg"

	"github.com/bakurits/fshare-common/auth"
	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// AuthorizeOptions options for authorization
type AuthorizeOptions struct {
	email string
}

// AuthorizeCommand stores dependencies for authorize command
type AuthorizeCommand struct {
	Host      string
	TokenPath string
}

// New : authorizeCmd represents the authorize command
func (a AuthorizeCommand) New() *cobra.Command {
	var opts AuthorizeOptions
	var authorizeCmd = &cobra.Command{
		Use:   "auth",
		Short: "make authorization in google drive with credentials with given directory which holds a file credentialsMail.json",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.authorize(opts)
		},
	}

	authorizeCmd.Flags().StringVarP(&opts.email, "email", "m", "", "email")
	_ = authorizeCmd.MarkFlagRequired("email")

	return authorizeCmd
}

// storeToken : make a get request to a server for getting token
func (a AuthorizeCommand) storeToken(email string, password string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, a.Host+cfg.GetTokenEndpoint, nil)
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

	if err := auth.SaveToken(a.TokenPath, &tok); err != nil {
		return err
	}
	return nil
}

// authorize : make authorization
func (a AuthorizeCommand) authorize(opts AuthorizeOptions) error {
	prompt := "Enter password:\n"
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	return a.storeToken(opts.email, password)
}
