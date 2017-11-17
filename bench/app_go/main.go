package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})
	fmt.Println("servindo em http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
