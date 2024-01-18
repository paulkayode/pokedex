package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
	"github.com/segunkayode1/pokedex/internal/cache"
	"time"
)

const interval = 5 * time.Second

var m_cache cache.Cache = cache.NewCache(interval)


func commandExplore(cfg *pokedexapi.Config, arg string) error{
    url := "https://pokeapi.co/api/v2/location-area/" + arg + "/"
	fmt.Printf("\nPokemon Located at %v: \n", arg)
	var httpBody []byte
	if val, ok := m_cache.Get(url); !ok{
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode > 399 {
			return fmt.Errorf("network-issue, status code: %v", resp.Status)
		}
		httpBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		m_cache.Put(url, httpBody)
   }else{
	    httpBody = val
   }
	locData := locationData{}
    err := json.Unmarshal(httpBody, &locData)
    if err != nil {
		return err
	}
	fmt.Println()
	for _,pokemonEncounter := range locData.PokemonEncounters{
		fmt.Println("\t- " + pokemonEncounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}


type locationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
