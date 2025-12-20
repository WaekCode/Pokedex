package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Result   []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type nextBack struct {
	NextURL     string
	PreviousURL string
}

func apiLocation(cfg *nextBack, prev bool) (location, error) {
	var url string
	if prev {
		url = cfg.PreviousURL
	} else {
		if cfg.NextURL == "" {
			url = "https://pokeapi.co/api/v2/location-area"
		} else {
			url = cfg.NextURL
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return location{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return location{}, fmt.Errorf("API request failed with status: %v", resp.Status)
	}

	var result location
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return location{}, err
	}

	cfg.NextURL = result.Next
	if result.Previous == nil {
		cfg.PreviousURL = ""
	} else {
		cfg.PreviousURL = *result.Previous
	}
	return result, nil
}