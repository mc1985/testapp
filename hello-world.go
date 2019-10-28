package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// Why no https?
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
