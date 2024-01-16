package pokedexapi

import (
	"io"
	"net/http"
	"errors"
)


func getJson(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	val, err:= io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 399 {
		return []byte{},errors.New("newtork error")
	}
	if err != nil {
		return []byte{}, err
	}

	return val, nil
	
}