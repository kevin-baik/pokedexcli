package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := os.Stdin
    scanner := bufio.NewScanner(reader)
    cliCommands := getCliCommands()
    for {
	if err := scanner.Err(); err != nil {
	    fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	fmt.Print("Pokedex > ")
	if scanner.Scan() {
	    text := scanner.Text()
	    texts := cleanInput(text)
	    command := texts[0]
	    commandInfo, ok := cliCommands[command]
	    if !ok {
		fmt.Println("Unknown command")
		continue
	    }
	    if err := commandInfo.callback(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	    }

	    
	}
    }
}
