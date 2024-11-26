package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

/*
	func Authorization(handler http.HandlerFunc, username, password []byte) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || user != string(username) || pass != string(password) {
				w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password"`)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("401 Unauthorized\n"))
				return
			}

			handler(w, r)
		}
	}
*/
func Authorization(next http.HandlerFunc) http.HandlerFunc {
	username := []byte("admin")
	password := []byte("123456")

	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "401 Unauthorized\n", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "401 Unauthorized\n", http.StatusUnauthorized)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			http.Error(w, "401 Unauthorized\n", http.StatusUnauthorized)
			return
		}

		credentials := string(decodedCredentials)
		parts := strings.Split(credentials, ":")
		if len(parts) != 2 {
			http.Error(w, "401 Unauthorized\n", http.StatusUnauthorized)
			return
		}

		if parts[0] != "userid" || parts[1] != "password" {
			http.Error(w, "401 Unauthorized\n", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "You have successfully accessed the /answer/ path!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "The answer is 42")
}

func main() {
	//username := []byte("admin")
	//password := []byte("123456")

	http.HandleFunc("/answer/", Authorization(answerHandler))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
