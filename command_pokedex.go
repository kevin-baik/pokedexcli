package main

import (
    "fmt"
)

func commandPokedex(cfg *config, arg string) error {
    listPokemon, err := cfg.pokeapiClient.ListPokedex()
    if err != nil {
	return err
    }
    fmt.Println("Your Pokedex:")
    totalPokemon := 0
    for _, pokemon := range listPokemon {
	totalPokemon++
	fmt.Println(" -", pokemon.Name)
    }
    fmt.Println("Total Pokemon:", totalPokemon)
    return nil
}
