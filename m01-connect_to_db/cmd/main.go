package main

import (
	"context"
	"fmt"
	"log"
	"m01-connect_to_db/config"
	"net/http"
)

func main() {
	config.ConnectToPostgresql()
	defer config.CloseConnection()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var greeting string
		err := config.Connection.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL!'").Scan(&greeting)
		if err != nil {
			http.Error(w, "Failed to query database", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, greeting)
	})
	log.Println("Starting server on : http://localhost:8765")
	err := http.ListenAndServe(":8765", nil)
	if err != nil {
		log.Fatal(err)
	}
}
