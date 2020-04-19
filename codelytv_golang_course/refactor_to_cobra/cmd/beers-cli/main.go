package main

import (
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/cli"
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {
	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
