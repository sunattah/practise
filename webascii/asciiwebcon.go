package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	g := banner
	
}
func main() {
	http.HandleFunc("/", aiwebHand)
	http.HandleFunc("/ascii-art", ascciiHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
