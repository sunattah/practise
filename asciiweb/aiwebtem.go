package main

import (
	"html/template"
	"net/http"
)

func aiwebHande(w http.ResponseWriter, r *http.Request) {
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
