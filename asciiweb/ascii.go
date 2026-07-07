package main

import (
	"fmt"
	"net/http"
)

func ascciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request occured", 405)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" {
		http.Error(w, "bad request", 400)
		return
	}
	if banner == "" {
		http.Error(w, "bad request", 400)
		return
	}
	fmt.Fprintf(w, "text %s, banner %s", text, banner)
}
func main() {
	http.HandleFunc("/", aiwebHand)
	http.HandleFunc("/ascii-art", ascciiHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
