package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	C "github.com/RemcoVeens/pokedex/commands"
	M "github.com/RemcoVeens/pokedex/models"
)

var commands = map[string]M.CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    C.CommandExit,
	},
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    C.CommandHelp,
	},
	"map": {
		Name:        "map",
		Description: "list towns",
		Callback:    C.CommandMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "list previus towns",
		Callback:    C.CommandMapB,
	},
	"explore": {
		Name:        "explore",
		Description: "fetch info a a area on the map",
		Callback:    C.CommandExplore,
	},
	"catch": {
		Name:        "catch",
		Description: "try to catch <pokemon>",
		Callback:    C.CommandCatch,
	},
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
			err = command.Callback(cleaned...)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
