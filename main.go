package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			fmt.Printf("Your command was: %v\n", cleaned[0])
		}
	}
}
