package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *nextBack, locationOrPokemon string,pokedex *map[string]PokemomDetails) error
}

func commandExit(cfg *nextBack, location string,pokedex *map[string]PokemomDetails) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *nextBack, location string,pokedex *map[string]PokemomDetails) error {

	res, err := apiLocation(cfg, false)
	if err != nil {
		fmt.Println("could not get locations")
		return err
	}

	for _, v := range res.Result {
		fmt.Println(v.Name)
	}

	return err
}

func commandExplore(cfg *nextBack, location string,pokedex *map[string]PokemomDetails) error {
	loc, err := locationDetails(cfg, location)
	if err != nil {
		fmt.Println("could not explore/find pokemones")
		return err
	}

	fmt.Printf("Exploring %v...\n", location)
	for _, v := range loc.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}

func commandMapb(cfg *nextBack, location string,pokedex *map[string]PokemomDetails) error {

	if cfg.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := apiLocation(cfg, true)


	for _, v := range res.Result {
		fmt.Println(v.Name)
	}

	return err
}

func commandCatch(cfg *nextBack, locationOrPokemon string,pokedex *map[string]PokemomDetails) error{
	res, err := getPokemoneDetails(cfg,locationOrPokemon)
	if err != nil {
		fmt.Println("could not get pokemone details")
		return err
	}


	if _, exists := (*pokedex)[res.Name]; exists {
		fmt.Printf("%v is already caught!\n", res.Name)
		return nil
	}

	// Catch logic
	

	baseExperience := res.BaseExperience                  // example base experience of the Pokémon
	catchThreshold := 255 - baseExperience // higher baseExperience → harder to catch
	randomNumber := rand.Intn(256) // 0–255


	fmt.Printf("Throwing a Pokeball at %v...\n", locationOrPokemon)


	if randomNumber < catchThreshold {
		fmt.Printf("%v was caught!!\n", res.Name)
		(*pokedex)[res.Name] = res
		fmt.Printf("Pokedex: %v\n", pokedex)
	} else {
		fmt.Printf("%v escaped!\n", res.Name)
	}

	return nil
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
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemone by name",
			callback:    commandCatch,
		},
	}
}

func commandHelp(cfg *nextBack, location string,pokedex *map[string]PokemomDetails) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range getCommands() {
		line := fmt.Sprintf("%v: %v", v.name, v.description)
		fmt.Println(line)
	}

	return nil

}
