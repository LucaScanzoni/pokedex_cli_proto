package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("no arguments allowed")
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.pokedex {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}