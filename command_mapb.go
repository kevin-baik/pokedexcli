package main

import (
    "io"
    "fmt"
    "net/http"
    "encoding/json"
)

func commandMapb(c *Config) error {
    if c.Previous == nil {
	return fmt.Errorf("you're on the first page")
    }
    resp, err := http.Get(*c.Previous)
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
    
    for _, location := range locationArea.Results {
	fmt.Println(location.Name)
    }

    c.Next = locationArea.Next
    c.Previous = locationArea.Previous

    return nil
}
