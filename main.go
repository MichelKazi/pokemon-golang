package main

import (
	"fmt"
	"log"

	"github.com/michelkazi/pokemon-golang/pokemon"
)

func main() {

	blaziken, err := pokemon.NewPokemon("blaziken")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blaziken)

	pikachu, err := pokemon.NewPokemon(25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pikachu)

	p, err := pokemon.NewPokemon(25.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
