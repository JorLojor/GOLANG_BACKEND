package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Println("Starting server on :8765")
	err := http.ListenAndServe(":8765", nil)
	if err != nil {
		log.Fatal(err)
	}
}
