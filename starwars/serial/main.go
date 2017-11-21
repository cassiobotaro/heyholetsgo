package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type character struct {
	Name      string `json:"name"`
	BirthYear string `json:"birth_year"`
}

func main() {
	nCharacters := 5
	for i := 1; i <= nCharacters; i++ {
		APIUrl := fmt.Sprintf("https://swapi.co/api/people/%d", i)
		fmt.Println(APIUrl)
		resp, err := http.Get(APIUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		char := character{}
		err = json.NewDecoder(resp.Body).Decode(&char)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s nasceu no ano %s.\n", char.Name, char.BirthYear)
	}
}
