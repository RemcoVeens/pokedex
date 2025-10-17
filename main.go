package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
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
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("could not exit")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func getCommand(input []string) (cliCommand, error) {
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
