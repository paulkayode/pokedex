package main
import(
	"fmt"
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
)
func commandPokedex(cfg *pokedexapi.Config, arg string) error{
	fmt.Println("Your Pokedex:")
	if(len(pokedex) == 0){
		fmt.Println("Is Empty")
		return nil
	}else{
		for name := range pokedex {
			fmt.Printf("\t- %v\n",name)
		}
		fmt.Println()
		return nil
	}
}