package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("You did not catch this pokemon yet")
		return nil
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, ty := range pokemon.Types {
		fmt.Printf("  - %v\n", ty.Type.Name)
	}
	return nil
}