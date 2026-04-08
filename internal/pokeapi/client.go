package pokeapi

import (
	"net/http"
	"time"

	"github.com/LucaScanzoni/pokedex_cli_proto/internal/pokecache"
)

// Client -
type Client struct {
	httpClient 	http.Client
	cache 		pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}
