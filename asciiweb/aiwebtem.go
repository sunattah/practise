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
	g := r.FormValue("banner")
	if g == "" {
		http.Error(w, "file is empty", 400)
		return
	}

	tem.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", aiwebHand)

	fmt.Println("server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
