package main

import (
	"fmt"
	"net/http"
)

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}

	if r.Header.Get("content-type") != "application/x-www-form-urlencoded" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	name := r.FormValue("username")
	language := r.FormValue("language")

	if name == "" {
		http.Error(w, "username required", 400)
		return
	}
	if language == "" {
		http.Error(w, "user language required", 400)
		return
	}
	fmt.Fprintf(w, "Hello %s, you are coding in %s!", name, language)

}

func main() {
	http.HandleFunc("/form", decodeHandler)

	fmt.Println("server runing on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
