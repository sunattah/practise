package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func aiwebHand(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tem, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}
	g := r.FormValue("banner")
	if g == "" {
		http.Error(w, "file is empty", 400)
		return
	}

	tem.Execute(w, nil)

}

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
