package main

import (
	"fmt"
	"net/http"
)

func aiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ASCII")
}

func main() {
	http.HandleFunc("/", aiHandler)

	fmt.Println("server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
} 