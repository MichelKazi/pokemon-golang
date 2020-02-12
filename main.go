package main

import (
	"fmt"
	"log"

	"github.com/michelkazi/pokemon-golang/pokemon"
)

func main() {
	p, err := pokemon.NewPokemon(25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
