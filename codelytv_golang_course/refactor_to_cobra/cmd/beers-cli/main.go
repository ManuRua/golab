package main

import (
	"flag"

	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/fetching"

	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/storage/ontario"

	beerscli "github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal"
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/cli"
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {
	// f, _ := os.Create("beers.cpu.prof")
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()

	// f, _ := os.Create("beers.mem.prof")
	// defer f.Close()
	// pprof.WriteHeapProfile(f)
}
