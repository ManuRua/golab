package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal"
	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/errors"
	jsoniter "github.com/json-iterator/go"
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

	err = b.betterUnmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}

func (b *beerRepo) standardUnmarshal(data []byte, beers *[]beerscli.Beer) error {
	err := json.Unmarshal(data, &beers)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return nil
}

func (b *beerRepo) betterUnmarshal(data []byte, beers *[]beerscli.Beer) error {
	var js = jsoniter.ConfigCompatibleWithStandardLibrary
	err := js.Unmarshal(data, &beers)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return nil
}
