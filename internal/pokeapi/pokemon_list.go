package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
    "fmt"
)

func (c *Client) ListPokemons(areaName string) (RespDeepLocations, error) {
    url := baseURL + "/location-area/" + areaName

    fmt.Println("url to explore:", url)
    
    if val, exists := c.cache.Get(url); exists {
	encountersInLocationResp := RespDeepLocations{}
	err := json.Unmarshal(val, &encountersInLocationResp)
	if err != nil {
	    return RespDeepLocations{}, err
	}
	return encountersInLocationResp, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
	return RespDeepLocations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
	return RespDeepLocations{}, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
	return RespDeepLocations{}, err
    }

    encountersInLocationResp := RespDeepLocations{}
    if err := json.Unmarshal(data, &encountersInLocationResp); err != nil {
	return RespDeepLocations{}, err
    }

    c.cache.Add(url, data)
    return encountersInLocationResp, nil
}
