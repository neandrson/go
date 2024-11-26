package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type Response struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

/*type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
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
}*/

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "ok",
		"result": map[string]string{
			"greetings": "hello",
			"name":      "stranger",
		},
	}
	json.NewEncoder(w).Encode(response)
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			output := Result{Greetings: "hello", Name: "stranger"}
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
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Name parameter is required", http.StatusBadRequest)
			return
		}
		if name == "admin" {
			panic("Invalid name")
		}
		next(w, r)
	}
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

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.Header().Set("Content-Type", "application/json")
				responseRpc := Response{}
				json.NewEncoder(w).Encode(responseRpc)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/rpc/", SetDefaultName(Sanitize(RPC(HelloHandler))))
	http.ListenAndServe(":8080", nil)
}
