package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
func NewPokemon(id interface{}) (*Pokemon, error) {
	var endpoint string

	// Check if the parameter passed is an int or string
	// This will determine what endpoint to hit when I query
	switch id.(type) {
	case string:
		endpoint = fmt.Sprintf("name/%s.json", id)
	case int:
		endpoint = fmt.Sprintf("id/%d.json", id)
	}

	// Configure the url now that we have an endpoint
	url := fmt.Sprintf("https://fizal.me/pokeapi/api/v2/%s", endpoint)

	// Making a GET request with our URL
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Deferring the Close() execution of res.Body is important to prevent resource leaks
	defer res.Body.Close()

	// Read the body of the response, and save it in var body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Get parsing boiz
	var data map[string]interface{}

	json.Unmarshal(body, &data)

	p := new(Pokemon)
	return p, nil
}
