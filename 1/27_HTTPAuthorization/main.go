package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

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
}
