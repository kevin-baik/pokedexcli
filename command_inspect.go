package main

import (
    "fmt"
)

func commandInspect(cfg *config, name string) error {
    // Check if pokemon in Pokedex
    pokemonResp, err := cfg.pokeapiClient.InspectPokemon(name)
    if err != nil {
	return err
    }
    pokemonStat := make(map[string]int)
    for _, stat := range pokemonResp.Stats {
	pokemonStat[stat.Stat.Name] = stat.BaseStat
    }
    fmt.Println("Name:", pokemonResp.Name)
    fmt.Println("Height:", pokemonResp.Height)
    fmt.Println("Weight:", pokemonResp.Weight)
    fmt.Println("Stats:")
    fmt.Println("\t-hp:", pokemonStat["hp"])
    fmt.Println("\t-attack:", pokemonStat["attack"])
    fmt.Println("\t-defense:", pokemonStat["defense"])
    fmt.Println("\t-special-attack:", pokemonStat["special-attack"])
    fmt.Println("\t-special-defense:", pokemonStat["special-defense"])
    fmt.Println("\t-speed:", pokemonStat["speed"])
    fmt.Println("Types:")
    for _, pType := range pokemonResp.Types {
	fmt.Println("\t-", pType.Type.Name)
    }
    return nil		
}
