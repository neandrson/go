package main

import (
	"fmt"
	"net/http"
)

/*type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type BasicAuthMiddleware struct {
	username string
	password string
}

func (m *BasicAuthMiddleware) Handle(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// проверяем формат заголовка авторизации
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Authorization required"`)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		username, password, err := parseAuthHeader(authHeader)
		if err != nil || username != m.username || password != m.password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Authorization required"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// если авторизация успешна, вызываем следующий обработчик
		next(w, r)
	}
}

type AuthenticatedHandler struct {
	authMiddleware MiddlewareFunc
	handlerFunc    HandlerFunc
}

func (ah *AuthenticatedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.authMiddleware(ah.handlerFunc)(w, r)
}

func answerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The answer is 42")
}

func parseAuthHeader(authHeader string) (string, string, error) {
	// удаляем префикс "Basic " из заголовка auth
	encodedCreds := strings.TrimPrefix(authHeader, "Basic ")

	// декодируем строку из base64
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return "", "", err
	}

	// cплитим логин и пароль по символу ":"
	creds := strings.SplitN(string(decodedCreds), ":", 2)
	if len(creds) != 2 {
		return "", "", fmt.Errorf("invalid auth header format")
	}

	return creds[0], creds[1], nil
}

func main() {
	authMiddleware := &BasicAuthMiddleware{
		username: "userid",
		password: "password",
	}

	answerHandler := &AuthenticatedHandler{
		authMiddleware: authMiddleware.Handle,
		handlerFunc:    answerHandlerFunc,
	}

	http.Handle("/answer/", answerHandler)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}*/

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			//http.Error(w, "Unauthorized", http.StatusUnauthorized)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username, password, ok := r.BasicAuth()
		if !ok || username != "John Doe" || password != "123456" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		/*if !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		credentials := string(decodedCredentials)
		parts := strings.Split(credentials, ":")
		if len(parts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if parts[0] != "userid" || parts[1] != "password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}*/

		next(w, r)
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The answer is 42"))
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
