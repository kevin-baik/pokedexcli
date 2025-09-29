package main

import (
    "io"
    "fmt"
    "net/http"
    "encoding/json"
)

const (
	locationAreaURL = "https://pokeapi.co/api/v2/location-area/"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func commandMap(c *Config) error {
    resp, err := http.Get(*c.Next)
    if err != nil {
	return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
	return err
    }
    
    var locationArea LocationAreas
    if err := json.Unmarshal(body, &locationArea); err != nil {
	return err
    }
    
    c.Next = locationArea.Next
    c.Previous = locationArea.Previous

    for _, location := range locationArea.Results {
	fmt.Println(location.Name)
    }

    return nil
}
