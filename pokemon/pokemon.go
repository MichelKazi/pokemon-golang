package pokemon

import (
	"fmt"
	"math"
	"strconv"
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
func NewPokemon(id int) *Pokemon {

	var endpoint string
	if math.IsNaN(id) {
		endpoint = id
	} else {
		endpoint = strconv.Itoa(id)
	}

	p := new(Pokemon)
	url := fmt.Sprintf("https://fizal.me/pokeapi/v2/%s/%s.json")
	return p
}
