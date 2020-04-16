package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"1": "Mad Jack Mixer",
	"2": "Keystone Ice",
	"3": "Belgian Moon",
}

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}
}
