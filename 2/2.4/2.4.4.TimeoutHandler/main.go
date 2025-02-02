package main

import (
	"io"
	"net/http"
	"time"
)

type TmpHandler struct{}

func (h TmpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, "http://localhost:8081/provideData", r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func StartServer(maxTimeout time.Duration) {
	http.Handle("/readSource", http.TimeoutHandler(TmpHandler{}, maxTimeout, ""))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}

func main() {
	StartServer(1 * time.Second)
}
