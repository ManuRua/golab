package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "stores-cli"}
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
