package main

import "fmt"



func commandHelp() error{
	fmt.Printf("\nThis is a cli tool for finding Information about pokemon!\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\t\t<command> [arguments]\n\n")
	fmt.Printf("The commands are:\n\n")

	for key, value := range getCommandMap(){
		fmt.Printf("\t\t%v: %v\n", key, value.description)
	}
	return nil
}
