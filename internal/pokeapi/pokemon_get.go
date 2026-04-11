package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name
	resp := RespPokemon{}
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return RespPokemon{}, err
		}
		return resp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}
	req_resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer req_resp.Body.Close()
	dat, err := io.ReadAll(req_resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}
	c.cache.Add(url, dat)
	err = json.Unmarshal(dat, &resp)
	if err != nil {
		return RespPokemon{}, err
	}
	return resp, nil
}