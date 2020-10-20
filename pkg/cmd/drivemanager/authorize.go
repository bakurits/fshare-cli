package drivemanager

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewAuthorizeCommand : authorizeCmd represents the authorize command
func NewAuthorizeCommand() *cobra.Command {
	var authorizeCmd = &cobra.Command{
		Use:   "authorize",
		Short: "make authorization in google drivemanager with credentials with given directory which holds a file credentials.json",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("too many arguments")
			}
			return authorize(args[0])
		},
	}
	return authorizeCmd
}

// authorize : make authorization
func authorize(filePath string) error {

	driveConfig := "state.json"

	credentialsMap := make(map[string]string)

	credentialsMap["credentialsPath"] = filePath
	fileJSON, _ := json.Marshal(credentialsMap)
	err := ioutil.WriteFile(driveConfig, fileJSON, 0777)

	return err
}
