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
func commandMap(cfg *pokedexapi.Config) error {
	res, err := pokedexapi.GetNextConfig(cfg)
	if err != nil {
		return err
	}
	printMap(res)
    *cfg = *res
	return nil
}

func commandMapb(cfg *pokedexapi.Config) error {
	res, err := pokedexapi.GetPrevConfig(cfg)
	if err != nil {
		return err
	}
	printMap(res)
    *cfg = *res
	return nil
}