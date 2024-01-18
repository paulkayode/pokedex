package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/segunkayode1/pokedex/internal/pokedexapi"
)


var pokedex map[string]pokemon = map[string]pokemon{}

func repl(cfg *pokedexapi.Config){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("pokedex > ")
	cliMap := getCommandMap()
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, " ")
		args := strings.Split(text, " ")
		text = args[0]
		if command,ok := cliMap[text]; ok {
			arg := ""
			if text == "explore" || text == "catch"  || text == "inspect"{
				if(len(args) == 2){
					arg = strings.ToLower(args[1])
				}else{
					fmt.Printf("Incorrect amount of arguments for %v 1 argument required\n",text)
					fmt.Printf("pokedex > ")
					continue
				}
			}

			err := command.callback(cfg,arg)
			if err != nil {
				fmt.Println(err)
			}
		}else{
				fmt.Println("command not found, input \"help\" to get list of available commands")
		}

		fmt.Printf("pokedex > ")
	}
}
type cliCom struct {
	name string
	description string
	callback func(cfg * pokedexapi.Config, arg string) error
}
func getCommandMap() map[string]cliCom {
	return map[string]cliCom {
		"help": {
			name: "help",
			description: "Displays a help Message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Lists 20 location areas in the Pokemon World, subsquent uses will list next 20",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists 20 location areas in the Pokemon World, subsequent uses will list previous 20",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Given name of location prints list of pokemon from that location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Given name of pokemon will attempt to catch pokemon if successfull, pokemon is added to pokedex",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Given name of pokemon, if previously caught displays information about pokemon",
			callback: commandInspect,
		},
	}
}


