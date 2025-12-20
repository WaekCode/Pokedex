package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	var s []string

	ful := ""

	for _, w := range text {
		if string(w) != " " {
			ful += string(w)
		} else {
			if ful != "" {
				s = append(s, strings.ToLower(ful)) // append the WORD
				ful = ""
			} // reset
		}
	}

	// append last word (VERY important)
	if ful != "" {
		s = append(s, strings.ToLower(ful))

	}
	return s
}


func startRepl() {
	cfg := &nextBack{}
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()

			Word := cleanInput(input)
			if len(Word) == 0 {
				continue
			}
			firstWord := Word[0]
			// fmt.Println("Your command was:",firstWord)

			cmd, ok := commands[firstWord]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				errs := cmd.callback(cfg)
				if errs != nil {
					fmt.Println(errs)
				}
			}

			err := scanner.Err()
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
		}

	}
}
