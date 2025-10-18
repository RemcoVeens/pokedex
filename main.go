package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type EncounterMethodRate struct{}
type NamedAPIResource struct{}
type PokemonEncounter struct{}

type location struct {
	id                     int32                 "json:id"
	name                   string                "json:name"
	game_index             int32                 "json:game_index"
	encounter_method_rates []EncounterMethodRate "json:encounter_method_rates"
	location               NamedAPIResource      "json:location"
	names                  []string              "json:names"
	pokemon_encounters     []PokemonEncounter    "json:pokemon_encounters"
}

var commands = map[string]models.cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "list towes",
		callback:    commandMap,
	},
}

const baseUrl = "https://pokeapi.co/api/v2/"

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap() error {
	url := fmt.Sprintf("%v%v", baseUrl, "")
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not read %v. %w", url, err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not body. %w", err)
	}
	locations := location{}
	if err := json.Unmarshal(body, &locations); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("locations: %v\n", locations)
	return nil
}

func getCommand(input []string) (models.cliCommand, error) {
	if len(input) < 1 {
		return cliCommand{}, fmt.Errorf("no command found")
	}
	command := input[0]
	for _, cmd := range commands {
		if command == cmd.name {
			return cmd, nil
		}
	}
	return cliCommand{}, fmt.Errorf("command not found")
}

func cleanInput(text string) (output []string) {
	for _, word := range strings.Split(text, " ") {
		if word != "" {
			output = append(output, strings.ToLower(word))
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			command, err := getCommand(cleaned)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			}
			command.callback()

		}
	}
}
