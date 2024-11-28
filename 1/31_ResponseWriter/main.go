package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	result := response{}
	result.Name = r.URL.Query().Get("name")
	jsonOut, _ := json.Marshal(result)
	log.Print(string([]byte(jsonOut)))
	json.NewEncoder(w).Encode(jsonOut)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
