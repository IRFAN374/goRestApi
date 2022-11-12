package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/allgrocery", AllGrocery).Methods("GET")
	router.HandleFunc("/grocery/{name}", GetGrocery).Methods("GET")
	router.HandleFunc("/grocery", AddGrocery).Methods("POST")
	router.HandleFunc("/grocery/{name}", UpdateGrocery).Methods("PUT")
	router.HandleFunc("/grocery/{name}", DeleteGrocery).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
