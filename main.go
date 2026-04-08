package main

import (
	"time"

	"github.com/LucaScanzoni/pokedex_cli_proto/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
