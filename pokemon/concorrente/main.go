package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	nPokemon := 151
	var wg sync.WaitGroup
	for i := 1; i <= nPokemon; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Downloading image %03d.png\n", i)
			resp, err := http.Get(fmt.Sprintf("https://assets.pokemon.com/assets/cms2/img/pokedex/full/%03d.png", i))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Salvando imagem %03d.png\n", i)
			file, err := os.Create(fmt.Sprintf("pokemon/%03d.png", i))
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(file, resp.Body)
			wg.Done()
			resp.Body.Close()
			file.Close()
		}(i)
	}
	wg.Wait()
}
