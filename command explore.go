package main

import (
	"encoding/json"
	"fmt"
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
)

func commandExplore(cfg *pokedexapi.Config, arg string) error{
    url := "https://pokeapi.co/api/v2/location-area/" + arg + "/"
	httpBody, err := pokedexapi.GetJson(url)
	if err != nil {
		return err
	}
	locData := locationData{}
    err = json.Unmarshal(httpBody, &locData)
    if err != nil {
		return err
	}
	fmt.Printf("\nExploring %v...\n", arg)
	fmt.Printf("\nFound Pokemon:\n\n")
	for _,pokemonEncounter := range locData.PokemonEncounters{
		fmt.Println(" - " + pokemonEncounter.Pokemon.Name)
	}
	fmt.Println()
	return nil
}


