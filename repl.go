package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"

    "github.com/kevin-baik/pokedexcli/internal/pokeapi"
    "github.com/kevin-baik/pokedexcli/internal/pokecache"
)

type config struct {
    pokeapiClient	pokeapi.Client
    pokeCache		pokecache.Cache
    nextLocationsURL	*string
    prevLocationsURL	*string
}

func startRepl(cfg *config) {
    reader := bufio.NewScanner(os.Stdin)
    for {
	if err := reader.Err(); err != nil {
	    fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	fmt.Print("Pokedex > ")
	reader.Scan()

	words := cleanInput(reader.Text())
	if len(words) == 0 {
	    continue 
	}

	commandName := words[0]
	var arg1 string
	if len(words) == 2 {
	    arg1 = words[1]
	}
	command, exists := getCommands()[commandName]
	if exists {
	    err := command.callback(cfg, arg1)
	    if err != nil {
		fmt.Println(err)
	    }
	    continue
	} else {
	    fmt.Println("Unknown command")
	    continue
	}
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

type cliCommand struct {
    name        string
    description string
    callback    func(*config, string) error
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
	"help": {
	    name:        "help",
	    description: "Displays a help message",
	    callback:	 commandHelp,
	},
	"map": {
	    name:        "map",
	    description: "Displays the next 20 locations",
	    callback:    commandMapf,
	},
	"mapb": {
	    name:        "mapb",
	    description: "Displays the previous 20 locations",
	    callback:    commandMapb,
	},
	"explore": {
	    name:        "explore",
	    description: "List of Pokemon in an area",
	    callback:    commandExplore,
	},
	"catch": {
	    name:        "catch",
	    description: "Attempt to catch a pokemon",
	    callback:    commandCatch,
	},
	"exit": {
	    name:        "exit",
	    description: "Exit the Pokedex",
	    callback:    commandExit,
	},
    }
}
