package main

import (
	"net/http"
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
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			//http.Error(w, "Unauthorized", http.StatusForbidden)
			//w.Write([]byte("Unauthorized"))
			//w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		/*auth := strings.SplitN(authHeader, " ", 2)
		//fmt.Println(auth)
		if len(auth) != 2 || auth[0] != "Basic" {
			//http.Error(w, "Unauthorized ", http.StatusUnauthorized)
			w.WriteHeader(http.StatusBadRequest)
			//w.Write([]byte("Unauthorized"))
			return
		}

		data, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			//http.Error(w, "Unauthorized ", http.StatusUnauthorized)
			w.WriteHeader(http.StatusBadRequest)
			//w.Write([]byte("Unauthorized"))
			return
		}*/
		//fmt.Println(string(data))
		/*if string(data) != "John:123456" {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}*/

		/*usernamePassword := strings.SplitN(string(data), ":", 2)
		if len(usernamePassword) != 2 || !checkCredentials(usernamePassword[0], usernamePassword[1]) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}*/
		/*u, p, ok := r.BasicAuth()
		fmt.Println(u, p)
		if !ok {
			//fmt.Println("Error parsing basic auth")
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized "))
			return
		}*/

		//fmt.Printf("Username: %s\n", u)
		//fmt.Printf("Password: %s\n", p)

		//w.WriteHeader(200)
		//next.ServeHTTP(w, r)
		return
	}
}

/*func checkCredentials(u, p string) bool {
	username := "John"
	password := "123456"
	// Здесь можно добавить логику проверки реальных учетных данных
	if username != u || password != p {
		//fmt.Printf("Username provided is correct: %s\n", u)
		return true
	}
	return false
}*/

func answerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "The answer is 42")
	w.Write([]byte("The answer is 42"))
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
