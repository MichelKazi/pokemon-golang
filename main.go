package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	// Configure the url now that we have an endpoint
	endpoint := "id/25.json"
	url := fmt.Sprintf("https://fizal.me/pokeapi/api/v2/%s", endpoint)

	// Making a GET request with our URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Deferring the Close() execution of res.Body is important to prevent resource leaks
	defer res.Body.Close()

	// Read the body of the response, and save it in var body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Get parsing boiz
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data["stats"].([]interface{})[5].(map[string]interface{})["base_stat"].(float64))

}
