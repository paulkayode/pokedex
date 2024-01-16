package main
import (
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
)


func main(){
	url := "https://pokeapi.co/api/v2/location-area"
	cfg := pokedexapi.Config{Next: &url}
	repl(&cfg)
}