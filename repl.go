package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func cleanInput(text string) []string{
	
	return strings.Fields(strings.ToLower(text))

}

func commandExit() error {
    fmt.Println ("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range existingCommand{
		
		fmt.Printf("%s: %s\n", command["name"], command["description"])
	}

    return nil
}
func REPLLoop(){

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for{
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())

		command, ok := existingCommand[input[0]]
		if !ok {
			fmt.Println("Unknown command")
		}else{
			err := command.callback()
 			if err != nil {
        			fmt.Printf("Error when %v\n", command)
 			}

		}

	}

}

