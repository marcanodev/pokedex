package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

type Payload struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name *string `json:"name"`
		Url  *string `json:"url"`
	} `json:"results"`
}

func GetJSON(url string) (JSON []byte, e error) {
	resp, err := http.Get(url)
	if err != nil {
		e = err
		return
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		err = errors.New("response failed")
	}

	if err != nil {
		e = err
		return
	}
	return body, e
}
