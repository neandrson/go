package main

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

type Result struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type Response struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

func (wr *wrappedResponseWriter) WriteHeader(code int) {
	wr.statusCode = code
	wr.ResponseWriter.WriteHeader(code)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.Context().Value("name").(string)

		if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) && name != "" {
			panic("invalid name format")
		}
		next(w, r)
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		ctx := context.WithValue(r.Context(), "name", name)
		next(w, r.WithContext(ctx))
	}
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				response := map[string]interface{}{
					"status": "error",
					"result": map[string]interface{}{},
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
				return
			}
		}()

		name := r.Context().Value("name").(string)

		response := map[string]interface{}{
			"status": "ok",
			"result": map[string]interface{}{
				"greetings": "hello",
				"name":      name,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	http.HandleFunc("/answer/", SetDefaultName(Sanitize(RPC(HelloHandler))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
