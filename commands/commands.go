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
var init_number int16

func CommandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	fmt.Println("help:          Displays a help message")
	fmt.Println("map:           List the 20 next area's")
	fmt.Println("mapb:          List the 20 previous area's")
	fmt.Println("explore(area): Lists pokemon in <area>")
	fmt.Println("exit:          Exit the Pokedex")
	return nil
}

func CommandMap(args ...string) error {
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
func CommandMapB(args ...string) error {
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

func CommandExplore(args ...string) error {
	// list's all local pokemons
	if len(args) == 1 {
		return fmt.Errorf("provide a argument of a town to explore")
	} else if len(args) > 2 {
		return fmt.Errorf("provide just one town to explore")
	}
	town := args[1]
	url := fmt.Sprintf("%v%v/%v/", baseUrl, "location-area", town)
	locations := M.LocationAreas{}
	val, found := Cache.Get(url)
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
	} else {
		if err := json.Unmarshal(val, &locations); err != nil {
			fmt.Println(err)
		}
	}
	for _, pokemon := range locations.Pokemon_encounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func CommandCatch(args ...string) error {
	if len(args) == 1 {
		return fmt.Errorf("provide a pokemon to catch")
	} else if len(args) > 2 {
		return fmt.Errorf("provide just one pokemon to catch")
	}
	pokemon := args[1]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	url := fmt.Sprintf("%v%v/%v/", baseUrl, "pokemon", pokemon)
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
	return nil
}
