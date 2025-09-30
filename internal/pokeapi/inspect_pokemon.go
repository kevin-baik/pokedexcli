package pokeapi

import (
    "fmt"
    "encoding/json"
)

func (c *Client) InspectPokemon(pokemon string) (RespPokemon, error) {
    // Check if pokemon exists in pokedex
    pokemonResp := RespPokemon{}
    pokeData, exists := c.Pokedex.Get(pokemon)
    if !exists {
	return RespPokemon{}, fmt.Errorf("%v is not in your Pokedex", pokemon)
    }
    err := json.Unmarshal(pokeData, &pokemonResp)
    if err != nil {
	return RespPokemon{}, err
    }
    return pokemonResp, nil
}
