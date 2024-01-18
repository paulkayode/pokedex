package main
import (
"github.com/segunkayode1/pokedex/internal/pokedexapi"
  "fmt"
)


func printMap(cfg *pokedexapi.Config){
	results := cfg.Results;
	for _, city := range results {
		fmt.Println(city.Name)
	}


}
func commandMap(cfg *pokedexapi.Config, arg string) error {
	if cfg.Next != nil && *cfg.Next == "https://pokeapi.co/api/v2" {
		*cfg.Next  = *cfg.Next + "/location-area"
	}
	res, err := pokedexapi.GetNextConfig(cfg)
	if err != nil {
		return err
	}
	printMap(res)
    *cfg = *res
	return nil
}

func commandMapb(cfg *pokedexapi.Config, arg string) error {
	res, err := pokedexapi.GetPrevConfig(cfg)
	if err != nil {
		return err
	}
	printMap(res)
    *cfg = *res
	return nil
}

