package main

import (
    "fmt"
)

func commandExplore(cfg *config, areaName string) error {
    fmt.Printf("Exploring %v...\n", areaName)
    
    encountersInLocationResp, err := cfg.pokeapiClient.ListPokemons(areaName)
    if err != nil {
	return err
    }

    for _, encounters := range encountersInLocationResp.PokemonEncounters {
	fmt.Println(encounters.Pokemon.Name)
    }

    return nil
}
