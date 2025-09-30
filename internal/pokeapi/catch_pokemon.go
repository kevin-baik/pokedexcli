package pokeapi

import (
    "net/http"
    "encoding/json"
    "io"
)

func (c *Client) GetPokemon(pokemon string) (RespPokemon, error) {
    url := baseURL + "/pokemon/" + pokemon

    if val, exists := c.cache.Get(url); exists {
	pokemonResp := RespPokemon{}
	err := json.Unmarshal(val, &pokemonResp)
	if err != nil {
	    return RespPokemon{}, err
	}
	return pokemonResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
	return RespPokemon{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
	return RespPokemon{}, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
	return RespPokemon{}, err
    }

    pokemonResp := RespPokemon{}
    if err := json.Unmarshal(data, &pokemonResp); err != nil {
	return RespPokemon{}, err
    }

    c.cache.Add(url, data)
    return pokemonResp, nil
}
