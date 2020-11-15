package drivemanager

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

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

// makeRequestBody : making request body
func makeRequestBody(email string, password string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return make([]byte, 0, 0), errors.New("can't make json object")
	}
	return requestBody, nil
}

// authorize : make authorization
func authorize(opts AuthorizeOptions) error {
	prompt := fmt.Sprintf("Enter password:\n")
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	fmt.Println(opts.email)
	fmt.Println(password)
	requestBody, err := makeRequestBody(opts.email, password)
	if err != nil {
		return err
	}
	fmt.Println(requestBody)

	http.Get("")
	return nil
}
