package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespSpecificLocation struct {
	Encounters  []Encounter	`json:"pokemon_encounters"`
	Name 		string 		`json:"name"`
}

type Encounter struct {
	Pokemon Pokemon	`json:"pokemon"`
}

type Pokemon struct {
	Name string	`json:"name"`
}