package pokeapi

type RespPokemon struct {
	Name	string		`json:"name"`
	BaseExp int			`json:"base_experience"`
	Height 	int			`json:"height"`
	Weight	int			`json:"weight"`
	Types 	[]struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	}					`json:"types"`
	Stats 	[]struct {
		Stat	struct {
			Name string `json:"name"`
		} 				`json:"stat"`
		Effort		int `json:"effort"`
		BaseStat	int	`json:"base_stat"`
	}					`json:"stats"`
}