package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/flights/prices", flightsHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Starting the server.....")
	fmt.Fprintf(w, "Requested Path : %q\n", r.URL.Path)
}

func flightsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Flight Resource : %q\n", r.URL.Path)
}
