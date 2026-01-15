package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string)   error {

	if len(cfg.caught) < 1 {
		fmt.Println("no pokemon caught")
	}else{
		for name, _ := range cfg.caught{
		  fmt.Printf("%s\n", name)
	  }
  }	
	return  nil
}


