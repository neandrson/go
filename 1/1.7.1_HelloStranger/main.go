package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "hello stranger")
		return
	}

	match, _ := regexp.MatchString("^[a-zA-Z]*$", name)
	if !match {
		fmt.Fprintf(w, "hello dirty hacker")
		return
	}

	fmt.Fprintf(w, "hello %s", name)
}

func main() {
	http.HandleFunc("/", StrangerHandler)
	http.ListenAndServe(":8080", nil)
}
