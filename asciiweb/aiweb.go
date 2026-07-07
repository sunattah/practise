package main

import (
	"fmt"
	"net/http"
)

func aiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ASCII")
}
