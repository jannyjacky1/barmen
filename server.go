package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func SomeHttpHandler(w http.ResponseWriter, req *http.Request) {
	url, _ := json.Marshal(req.URL)
	w.WriteHeader(200)
	w.Write(url)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", SomeHttpHandler)
	mux.HandleFunc("/hello", SomeHttpHandler)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
