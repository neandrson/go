package main

import (
	"fmt"
	"net/http"
	"strings"
)

func SetDefaultName(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			r.URL.Query().Set("name", "stranger")
		}
	})
}

func Sanitize(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		var dirty bool = false
		for _, n := range "0123456789_[]{}./,()`!@#$%^&*-+=\"'" {
			if strings.ContainsRune(name, n) {
				dirty = true
			}
		}
		if dirty {
			r.URL.Query().Set("name", "dirty hacker")
		}
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprint(w, "Hello, ", name)
}

func main() {
	sanitize := http.HandlerFunc(Sanitize(HelloHandler))
	//setDefaultName := http.Handler(sanitize)

	http.Handle("/", sanitize)

	http.ListenAndServe(":8080", nil)
}
