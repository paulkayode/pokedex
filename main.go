package main

import (
	"fmt"
	"bufio"
	"os"
)
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
	}
}


func commandHelp() error{
	fmt.Printf("\nThis is a cli tool for finding Information about pokemon!\n\n")
	fmt.Printf("Usage\n\n")
	fmt.Printf("\t\t<command> [arguments]\n\n")
	fmt.Printf("The commands are:\n\n")

	for _, value := range getCommandMap(){
		fmt.Printf("\t\t%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandExit() error{
	os.Exit(0)
	return nil
}

func main(){
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