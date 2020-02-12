package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Pokemon struct
type Pokemon struct {
	name      string
	id        int
	hp        int
	attack    int
	defense   int
	weight    int
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

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// And now that we finally parsed the data, we can assign values to our pokemon
	p := new(Pokemon)
	p.name = data["name"].(string) // The .(string) is type assertion
	p.id = data["id"].(int)        // Type assertion provides access to the actual value of an interface

	// Getting into the embedded fields of a JSON is actually really, REALLY tricky though
	// This means we have to know the types of each embedded JSON object
	// In this case, data["stats"] needs to be asserted to ([]interface{}) because it's a slice we need to index
	// Then, data["stats"].([]interface)[5] needs to be asserted to (map[string]interface{}) a map of strings to interfaces
	// Finally, the ["base_stat"] key is asserted to a (float64), which we'll ultimately convert to an int
	p.hp = int(data["stats"].([]interface{})[5].(map[string]interface{})["base_stat"].(float64))      // From here on it's pretty simple to access keys, indices and values
	p.attack = int(data["stats"].([]interface{})[4].(map[string]interface{})["base_stat"].(float64))  // From here on it's pretty simple to access keys, indices and values
	p.defense = int(data["stats"].([]interface{})[3].(map[string]interface{})["base_stat"].(float64)) // From here on it's pretty simple to access keys, indices and values
	return p, nil
}
