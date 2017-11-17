package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

var comments []string

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == http.MethodPost {
		comments = append(comments, r.FormValue("comentario"))
	}
	form := `
 <form action="" method="post">
	<p><input type=text name=comentario>
</form>
	`
	body := strings.Join(comments, "<br/>") + form
	fmt.Fprint(w, body)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/comments", commentsHandler)
	fmt.Println("Servindo em http://localhost:8000")
	port := os.Getenv("PORT") || "8000"
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
