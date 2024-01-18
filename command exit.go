package main

import ("os"
"github.com/segunkayode1/pokedex/internal/pokedexapi"
)
func commandExit(cfg *pokedexapi.Config, arg string) error{
	os.Exit(0)
	return nil
}