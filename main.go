package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	M "github.com/RemcoVeens/pokedex/models"
)

var commands = map[string]M.CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	},
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp,
	},
	"map": {
		Name:        "map",
		Description: "list towes",
		Callback:    commandMap,
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
	locations := M.Location{}
	if err := json.Unmarshal(body, &locations); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("locations: %v\n", locations)
	return nil
}

func getCommand(input []string) (M.CliCommand, error) {
	if len(input) < 1 {
		return M.CliCommand{}, fmt.Errorf("no command found")
	}
	command := input[0]
	for _, cmd := range commands {
		if command == cmd.Name {
			return cmd, nil
		}
	}
	return M.CliCommand{}, fmt.Errorf("command not found")
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
			command.Callback()

		}
	}
}
