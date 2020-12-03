package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdCompletion() *cobra.Command {
	var shellType string
	var filename string

	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion scripts",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			rootCmd := cmd.Parent()
			w, _ := os.Create(filename)
			switch shellType {
			case "bash":
				return rootCmd.GenBashCompletion(w)
			case "zsh":
				return rootCmd.GenZshCompletion(w)
			case "powershell":
				return rootCmd.GenPowerShellCompletion(w)
			case "fish":
				return rootCmd.GenFishCompletion(w, true)
			default:
				return fmt.Errorf("unsupported shell type %q", shellType)
			}
		},
	}

	cmd.Flags().StringVarP(&shellType, "shell", "s", "", "Shell type: {bash|zsh|fish|powershell}")
	cmd.Flags().StringVarP(&filename, "filename", "f", "", "")

	return cmd
}
