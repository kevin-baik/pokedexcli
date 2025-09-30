package pokeapi

import (
    "encoding/json"
)

func (c *Client) ListPokedex() ([]RespPokemon, error) {
    listPokemon, err := c.Pokedex.ListAll()
    if err != nil {
	return []RespPokemon{}, err
    }

    listPokedex := []RespPokemon{}
    for _, pokemon := range listPokemon {
	pokemonResp := RespPokemon{}
	err := json.Unmarshal(pokemon, &pokemonResp)
	if err != nil {
	    return []RespPokemon{}, err
	}
	listPokedex = append(listPokedex, pokemonResp)
    }
    return listPokedex, nil
}
