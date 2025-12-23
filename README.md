# Pokedex 🧭

**A simple command-line Pokedex written in Go** that uses the public PokeAPI (https://pokeapi.co/) to list locations, explore encounters, and attempt to catch Pokémon (with a simple random-based capture mechanic). Responses from the API are cached in-memory to reduce network usage.

---

## Features ✅

- Interactive REPL-style CLI (`Pokedex >`) with the following commands:
  - `map` — list available location areas from the PokeAPI
  - `mapb` — show previous page of locations
  - `explore <location>` — list Pokémon encounters for a location
  - `catch <pokemon>` — attempt to catch a Pokémon by name
  - `pokedex` — list all Pokémon you've caught
  - `inspect <pokemon>` — show details (height, weight, stats, types) of a caught Pokémon
  - `help` — show available commands
  - `exit` — quit the program

- Simple in-memory cache for API responses (configurable expiration)
- Lightweight codebase with focused tests (input cleaning, cache behavior, location cache usage)

---

## Prerequisites ⚙️

- Go 1.20+ (or your installed Go toolchain)
- Internet access to query https://pokeapi.co/

