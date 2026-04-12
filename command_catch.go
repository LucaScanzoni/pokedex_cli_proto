package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ... string) error {
	if len (args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	random := rand.Intn(pokemon.BaseExp)
	if random < 40 {
		cfg.pokedex [name] = pokemon
		fmt.Printf("%v was caught!\n", name)
	} else {
		fmt.Printf("%v escaped\n", name)
	}
	return nil
}	