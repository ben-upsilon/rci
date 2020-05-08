package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	httpServer()
}
func testCheckJava() {
	cmd := exec.Command("java", "--version")
	cmd.Dir = ""
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("wtf")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func httpServer() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Print("server run by http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
