package main

import (
	"fmt"
	"net/http"
)

func aiwebHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Hello ASCII")
		return

	}

}
