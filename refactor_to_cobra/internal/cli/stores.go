package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"1": "Mercadona",
	"2": "Carrefour",
	"3": "Alcampo",
}

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
