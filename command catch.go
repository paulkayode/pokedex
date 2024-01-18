package main
import (
	"encoding/json"
	"fmt"
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
	"math/rand"
	"time"
)

func commandCatch(cfg *pokedexapi.Config, arg string) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	url :=  "https://pokeapi.co/api/v2/pokemon/" + arg + "/"

	fmt.Printf("Throwing a Pokeball at %v...\n", arg)

	if _,ok := pokedex[arg]; ok{
		fmt.Printf("%v has already been caught!\n", arg)
		return nil
	}

	httpBody, err := pokedexapi.GetJson(url)
	if err != nil {
		return err
	}
	pokMon := pokemon{}
	err = json.Unmarshal(httpBody, &pokMon)
	if err != nil {
		return err
	}

	val := int32(pokMon.BaseExperience)

	if  random.Int31n(800) >= val {
		pokedex[arg] = pokMon
		fmt.Printf("%v was caught!\n", arg)
	}else {
		fmt.Printf("%v escaped!\n", arg)
	}
	
    return nil
}