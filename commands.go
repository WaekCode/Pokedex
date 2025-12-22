package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *nextBack,location string) error
}

func commandExit(cfg *nextBack, location string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *nextBack, location string) error {

	res, err := apiLocation(cfg, false)

	// fmt.Println(res.Next)
	// fmt.Println(res.Previous)

	for _, v := range res.Result {
		fmt.Println(v.Name)
	}

	return err	
}

func commandExplore(cfg *nextBack, location string) error {
	loc, err := locationDetails(cfg,location)
	if err != nil {
		fmt.Println("could not explore/find pokemones")
	}

	fmt.Printf("Exploring %v...\n", location)
	for _, v := range loc.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}

func commandMapb(cfg *nextBack, location string) error {

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
		"explore": {
			name:        "explore",
			description: "Explore a location and see which pokemones are there",
			callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *nextBack, location string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range getCommands() {
		line := fmt.Sprintf("%v: %v", v.name, v.description)
		fmt.Println(line)
	}

	return nil

}
