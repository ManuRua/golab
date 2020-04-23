package fetching

import (
	beerscli "github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal"
	"github.com/pkg/errors"
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

	for _, beer := range beers {
		if beer.ProductID == id {
			return beer, nil
		}
	}

	return beerscli.Beer{}, errors.Errorf("Beer %d not found", id)
}
