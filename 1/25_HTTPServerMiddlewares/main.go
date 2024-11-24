package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if isDirty, _ := regexp.MatchString("[^a-zA-Z]", name); isDirty {
			fmt.Fprint(w, "Hello, dirty hacker!")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			r.URL.RawQuery = "name=stranger"
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	hellohandler := http.HandlerFunc(HelloHandler)
	//setDefaultName := http.Handler(sanitize)
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloHandler)

	sanitize := http.Handle("/", Sanitize(hellohandler))
	handler := Sanitize(SetDefaultName(sanitize))

	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
