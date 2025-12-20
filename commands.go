package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*nextBack) error

}


func commandExit(cfg *nextBack) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *nextBack) error {
	
	res, err := apiLocation(cfg, false)

	// fmt.Println(res.Next)
	// fmt.Println(res.Previous)

	for _, v := range res.Result {
		fmt.Println(v.Name)
	}

	return err
}


func commandMapb(cfg *nextBack) error {
	
	if cfg.PreviousURL == "" {
        fmt.Println("you're on the first page")
        return nil
    }
	
	res, err := apiLocation(cfg, true)

	// fmt.Println(res.Next)
	// fmt.Println(res.Previous)

	for _, v := range res.Result {
		fmt.Println(v.Name)
	}

	return err
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays locations from the PokeAPI",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous locations from the PokeAPI",
			callback:    commandMapb,
		},
	}
}

func commandHelp(cfg *nextBack) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range getCommands() {
		line := fmt.Sprintf("%v: %v", v.name, v.description)
		fmt.Println(line)
	}

	return nil

}