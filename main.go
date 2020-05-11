package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	//httpServer()

	CloneByTag()
}

func httpServer() {
	http.HandleFunc("/status", CheckTaskStatus)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Print("server run by http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
