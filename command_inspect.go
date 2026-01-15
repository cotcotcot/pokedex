package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string)   error {
	if len(args) < 1 {
		fmt.Println("Usage: inspect pokemonName")
	}

	name := args[0]
	
	pokemon, exist:= cfg.caught[name]
	if exist{
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats :")

	}else{
		fmt.Println("This pokemon is not in your bag") 
	}

	
	return  nil
}


