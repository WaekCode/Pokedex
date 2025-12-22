package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

type PokemomDetails struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	BaseExperience int `json:"base_experience"`
	
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}


func getPokemoneDetails(cfg *nextBack, pokemon string) (PokemomDetails, error) {
	if pokemon != ""{
	pokemonurl := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"
	var p PokemomDetails

	c := cfg.Cache
	res,ok := c.Get(pokemonurl)
	if ok{
		fmt.Println("Cache resource being used...")
		err1 := json.Unmarshal(res, &p)
		if err1 != nil{
			return PokemomDetails{},err1
		}
		return p,nil
	}


	req, err2 := http.NewRequest("GET",pokemonurl,nil)
	if err2 != nil{
		return  PokemomDetails{},err2
	}


	// client
	client := &http.Client{}
	resp,err3 := client.Do(req)

	if err3 != nil{
		return PokemomDetails{},err3
	}

	defer resp.Body.Close()

	respBytes, err4 := io.ReadAll(resp.Body)
	if err4 != nil {
		return PokemomDetails{}, err4
	}

	c.Add(pokemonurl,respBytes)

	if err5 := json.Unmarshal(respBytes, &p); err5 != nil {
		return PokemomDetails{}, err5
	}
	return p,nil}
	return PokemomDetails{},fmt.Errorf("pokemon name cannot be empty")

}