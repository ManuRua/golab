package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal"
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/errors"
)

const (
	productsEndpoint = "/products"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type beerRepo struct {
	url string
}

// NewOntarioRepository fetch data from http
func NewOntarioRepository() beerscli.BeerRepo {
	return &beerRepo{
		url: ontarioURL,
	}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEndpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", productsEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", productsEndpoint)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}
