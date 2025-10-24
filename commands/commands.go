package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	I "github.com/RemcoVeens/pokedex/internal"
	M "github.com/RemcoVeens/pokedex/models"
)

const baseUrl = "https://pokeapi.co/api/v2/"

var Cache = I.NewCache(5 * time.Second)

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

var init_number int16

func CommandMap() error {
	for i := init_number; i < init_number+20; i++ {
		url := fmt.Sprintf("%v%v/%d/", baseUrl, "location-area", i+1)
		val, found := Cache.Get(url)
		locations := M.LocationAreas{}
		if !found {
			res, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("could not read %v. %w", url, err)
			}
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return fmt.Errorf("could not body. %w", err)
			}
			if err := json.Unmarshal(body, &locations); err != nil {
				fmt.Println(err)
			}
			raw, err := json.Marshal(locations)
			Cache.Add(url, raw)
		} else {
			if err := json.Unmarshal(val, &locations); err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("%v\n", locations.Name)
	}
	init_number += 20
	return nil
}
func CommandMapB() error {
	if init_number > 40 {
		init_number -= 20
	} else {
		init_number = 0
	}
	err := CommandMap()
	if err != nil {
		return err
	}
	return nil
}
