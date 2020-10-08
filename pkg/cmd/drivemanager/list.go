package drivemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bakurits/fileshare/pkg/drive"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// NewListCommand : generates of command list command
func NewListCommand() *cobra.Command {

	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "show list of files",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runList()
			return err
		},
	}

	return listCmd
}

func runList() error {
	content, err := ioutil.ReadFile("state.json")

	if err != nil {
		return errors.New("you are not authorized")
	}

	var credentialsMap map[string]string

	_ = json.Unmarshal(content, &credentialsMap)

	credentialsPath, ok := credentialsMap["credentialsPath"]
	if !ok {
		return errors.New("you are not authorized")
	}

	service, err := drive.Authorize(credentialsPath)

	if err != nil {
		return err
	}

	list := service.List()

	for i := range list {
		fmt.Println(list[i].Name)
	}

	return nil
}
