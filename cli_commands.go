package main

import (
    "fmt"
)

type cliCommand struct {
    name	string
    description	string
    callback	func() error
}

func getCliCommands() map[string]cliCommand {
    return map[string]cliCommand {
	"exit": {
	    name:		"exit",
	    description:	"Exit the Pokedex",
	    callback:	commandExit,
	},
	"help": {
	    name:		"help",
	    description:	"Displays a help message",
	    callback:	commandHelp,
	},
    }
}

func commandExit() error {
    return fmt.Errorf("Closing the Pokedex... Goodbye!")
    
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")
    for _, command := range getCliCommands() {
	fmt.Printf("%v: %v\n", command.name, command.description)
    }
    return nil
}
