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

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := Response{Status: "ok", Result: Result{Greetings: "hello", Name: name}}
	json.NewEncoder(w).Encode(response)
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				// Возвращаем ответ с ошибкой
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				output := Response{Status: "error", Result: Result{}}
				jsonOutput, err := json.Marshal(output)
				if err != nil {
					return
				}
				json.NewEncoder(w).Encode(jsonOutput)
				//w.Write(jsonOutput)
				return
			}

		}()

		//name := r.Context().Value("name").(string)
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		response := Response{
			Status: "ok",
			Result: Result{
				Greetings: "hello",
				Name:      name,
			},
		}
		//w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		//next.ServeHTTP(w, r)
	}
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

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				// Возвращаем ответ с ошибкой
				//w.Header().Set("Content-Type", "application/json")
				//w.WriteHeader(200)
				//output := Response{Status: "error", Result: Result{}}
				/*jsonOutput, err := json.Marshal(output)
				if err != nil {
					return
				}*/
				jsonOutput := `{"status":"error","result":{}}`
				//json.NewEncoder(w).Encode(jsonOutput)
				w.Write([]byte(jsonOutput))
				//next.ServeHTTP(w, r)
				return
			}

		}()
		name := r.URL.Query().Get("name")
		var dirty bool = false
		for _, n := range "0123456789_[]{}./,()`!@#$%^&*-+=\"'" {
			if strings.ContainsRune(name, n) {
				dirty = true
				break
			}
		}
		if dirty {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)
	})
}

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	output := Result{Greetings: "hello", Name: name}
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		return
	}
	w.Write(jsonOutput)
}

func main() {
	http.HandleFunc("/rpc/", SetDefaultName(Sanitize(RPC(HelloHandler))))
	http.ListenAndServe(":8080", nil)
}
