package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io"
	"github.com/WaekCode/Pokedex/internal/pokecache"
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
	Cache       *pokecache.Cache
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

	c := cfg.Cache
	res, ok := c.Get(url)

	var result location
	if ok {
		fmt.Println("Cache resource being used...")
		if err := json.Unmarshal(res, &result); err != nil {
			return location{}, err
		}
	} else {
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

		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return location{}, err
		}

		// Add to cache
		c.Add(url, respBytes)

		// Decode for returning
		if err := json.Unmarshal(respBytes, &result); err != nil {
			return location{}, err
		}
	}

	// Update next/previous URLs regardless of source (cache or API)
	cfg.NextURL = result.Next
	if result.Previous == nil {
		cfg.PreviousURL = ""
	} else {
		cfg.PreviousURL = *result.Previous
	}

	return result, nil
}
