package main

import (
	"encoding/base64"
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
			http.Error(w, "Authorization required", http.StatusForbidden)
			return
		}

		auth := strings.SplitN(authHeader, " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Invalid authorization", http.StatusUnauthorized)
			return
		}

		data, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			http.Error(w, "The answer is 42", http.StatusUnauthorized)
			return
		}

		if string(data) != "username:password" {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "The answer is 42")
	w.Write([]byte("The answer is 42\n"))
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

/*func main() {
	//username := []byte("admin")
	//password := []byte("123456")

	http.HandleFunc("/answer/", Authorization(answerHandler))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}*/
