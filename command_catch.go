package main

import (
    "fmt"
    "math/rand/v2"
    "encoding/json"
)

func commandCatch(cfg *config, arg string) error {
    fmt.Printf("Throwing a Pokeball at %v...\n", arg)

    pokemonResp, err := cfg.pokeapiClient.GetPokemon(arg)
    if err != nil {
	return err
    }

    baseExp := pokemonResp.BaseExperience
    randFloat := float64(rand.IntN(100))
    chance := randFloat / float64(baseExp)
    
    fmt.Println("randFloat:", randFloat)
    fmt.Println("baseExp:", baseExp)
    fmt.Println("My Chances:", chance)
    
    if chance > 0.50 {
	jsonData, err := json.Marshal(pokemonResp)
	if err != nil {
	    return err
	}
	cfg.pokeapiClient.Pokedex.Add(arg, jsonData)	
	fmt.Printf("%v was caught!", arg)
    } else {
	fmt.Printf("%v escaped!\n", arg)
    }
    return nil
}
