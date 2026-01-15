package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("Usage: explore location")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil{
		return err
	}

	fmt.Printf("Exploring %s...\n", name)

	pokelist := location.PokemonEncounters

	if len(pokelist) == 0{
		fmt.Println("No pokemon found.")
	}else{

		for _, encounter := range pokelist{
			fmt.Printf("%s\n", encounter.Pokemon.Name)
		}
	}

	return nil
}

