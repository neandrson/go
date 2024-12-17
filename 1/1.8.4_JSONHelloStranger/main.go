package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Result struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type Response struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

var dirty bool = false

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := Response{Status: "ok", Result: Result{Greetings: "hello", Name: name}}
	json.NewEncoder(w).Encode(response)

}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		defer func() {
			if recover() != nil {
				// Возвращаем ответ с ошибкой
				//w.Header().Set("Content-Type", "application/json")
				//w.WriteHeader(200)
				//output := Response{Status: "error", Result: Result{}}
				//jsonOutput, err := json.Marshal(output)
				//if err != nil {
				//	return
				//}
				//jsonOutput := `{"status":"error","result":{}}`
				//json.NewEncoder(w).Encode(jsonOutput)
				//w.Write([]byte(name))
				next.ServeHTTP(w, r)
				return
			}
		}()

		if name == "" {
			next.ServeHTTP(w, r)
			return
		}

		for _, n := range "0123456789_[]{}./,()`!@#$%^&*-+=\"'" {
			if strings.ContainsRune(name, n) {
				dirty = true
			}
		}
		if dirty {
			panic("Invalid name")
		}
		//w.Write([]byte(name))
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			output := Response{Status: "ok", Result: Result{Greetings: "hello", Name: "stranger"}}
			jsonOutput, err := json.Marshal(output)
			if err != nil {
				return
			}
			w.Write(jsonOutput)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		defer func() {
			if recover() != nil {
				// Возвращаем ответ с ошибкой
				//w.Header().Set("Content-Type", "application/json")
				//w.WriteHeader(200)
				//output := Response{Status: "error", Result: Result{}}
				//jsonOutput, err := json.Marshal(output)
				jsonOutput := `{"status":"error","result":{}}`
				//if err != nil {
				//	return
				//}
				//json.NewEncoder(w).Encode(jsonOutput)
				w.Write([]byte(jsonOutput))
				return
			}

		}()
		if name == "" {
			next.ServeHTTP(w, r)
			return
		}
		for _, n := range "0123456789_[]{}./,()`!@#$%^&*-+=\"'" {
			if strings.ContainsRune(name, n) {
				dirty = true
			}
		}
		if dirty {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/rpc/", RPC(SetDefaultName(Sanitize(HelloHandler))))
	http.ListenAndServe(":8080", nil)
}
