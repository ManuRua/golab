package main

import (
	"github.com/ManuRua/golab/refactor_to_cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
