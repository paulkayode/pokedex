package main

import (
	"fmt"
	"bufio"
	"os"
)


func repl(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("pokedex > ")
	cliMap := getCommandMap()
	for scanner.Scan() {
		text := scanner.Text()
		if _,ok := cliMap[text]; ok {
			cliMap[text].callback()
		}

		fmt.Printf("pokedex > ")
	}
}
type cliCom struct {
	name string
	description string
	callback func() error
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
	}
}