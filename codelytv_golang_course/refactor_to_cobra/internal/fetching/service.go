package fetching

import (
	"math"
	"sync"

	beerscli "github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal"
)

type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]beerscli.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (beerscli.Beer, error)
}

type service struct {
	beerRepo beerscli.BeerRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

func (s *service) FetchBeers() ([]beerscli.Beer, error) {
	return s.beerRepo.GetBeers()
}

func (s *service) FetchByID(id int) (beerscli.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return beerscli.Beer{}, err
	}

	beersPerRoutine := 10
	numRoutines := numOfRoutines(len(beers), beersPerRoutine)

	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)

	var b beerscli.Beer

	for i := 0; i < numRoutines; i++ {
		func(id, begin, end int, beers []beerscli.Beer, b *beerscli.Beer, wg *sync.WaitGroup) {
			for i := begin; i <= end; i++ {
				if beers[i].ProductID == id {
					*b = beers[i]
				}
			}
			wg.Done()
		}(id, i, i+beersPerRoutine, beers, &b, wg)
	}

	wg.Wait()

	return b, nil
}

func numOfRoutines(numOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numOfBeers) / float64(beersPerRoutine)))
}
