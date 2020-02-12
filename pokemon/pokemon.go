package pokemon

import (
	"fmt"
)

// Pokemon struct
type Pokemon struct {
	id        int8
	hp        int8
	attack    int8
	defense   int8
	weight    int8
	spriteURL string
	abilities []string
}

// NewPokemon is a constructor for pokemon
func NewPokemon(id interface{}) *Pokemon {
	var endpoint string

	// Check if the parameter passed is an int or string
	// This will determine what endpoint to hit when I query
	switch id.(type) {
	case string:
		endpoint = fmt.Sprintf("name/%s.json", id)
	case int:
		endpoint = fmt.Sprintf("id/%d.json", id)
	}

	p := new(Pokemon)
	url := fmt.Sprintf("https://fizal.me/pokeapi/v2/%s", endpoint)

	return p
}
