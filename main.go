package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello bosq")
}

func handleRequests() {
	// using gorilla mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("App running on port 8081")

	handleRequests()
}
