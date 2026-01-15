package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string)   error {
	if len(args) < 1 {
		return errors.New("Usage: catch pokemonName")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil{
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	possibilities := []string{
		"Yes",
		"No"}
	answers:= possibilities[rand.Intn(len(possibilities))] 
	if answers == "Yes" {
		cfg.caught[name] = pokemon
		fmt.Printf("%s est pris\n", name)
	}else{
		fmt.Printf("%s s’est échappé\n", name)
	}
	return  nil
}


