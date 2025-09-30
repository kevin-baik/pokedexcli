package main

import (
    "fmt"
)

func commandExplore(cfg *config, arg string) error {
    fmt.Printf("Exploring %v...\n", arg)
    
    encountersInLocationResp, err := cfg.pokeapiClient.ListPokemons(arg)
    if err != nil {
	return err
    }

    for _, encounters := range encountersInLocationResp.PokemonEncounters {
	fmt.Println(encounters.Pokemon.Name)
    }

    return nil
}
