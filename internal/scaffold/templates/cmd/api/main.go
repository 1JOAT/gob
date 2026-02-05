package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to your new gob project!")
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
