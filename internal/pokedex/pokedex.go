package pokedex

import (
	"github.com/mjossany/Gokedex/internal/pokeapi"
)

type Pokedex struct {
	Pokedex map[string]pokeapi.RespPokemonInfo
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokedex: make(map[string]pokeapi.RespPokemonInfo),
	}
}
