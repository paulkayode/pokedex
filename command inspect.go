package main

import ("fmt"
"github.com/segunkayode1/pokedex/internal/pokedexapi"
)
func commandInspect(cfg *pokedexapi.Config, arg string) error{
	if val, ok := pokedex[arg]; !ok {
		fmt.Printf("You have not caught %v yet\n", arg)
       return nil
	}else {
		fmt.Printf("\nName: %v\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Printf("Stats:\n")
		for _,stat := range val.Stats {
			fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _,poktype := range val.Types {
			fmt.Printf("\t-%v\n", poktype.Type.Name)
		}
		fmt.Println()
		return nil
	}
}