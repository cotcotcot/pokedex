package main

import (
)

var 	existingCommand = map[string]cliCommand{
    		"exit": {
        		name:        "exit",
        		description: "Exit the Pokedex",
        		callback:    commandExit,
    		},
    		"help": {
	    	name: "help",
	    	description: "Exit the Pokedex",
	    	callback:    commandHelp,
    		},
	}	
func main(){
	REPLLoop()
	return
}


