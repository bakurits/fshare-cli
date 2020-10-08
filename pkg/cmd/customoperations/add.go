package customoperations

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// NewAddCommand :  returns sum of numbers
func NewAddCommand() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "add list of numbers",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			sum := 0
			for _, args := range args {

				num, err := strconv.Atoi(args)

				if err != nil {
					fmt.Println(err)
				}
				sum = sum + num
			}
			fmt.Println("result of addition is", sum)
		},
	}
	return addCmd
}
