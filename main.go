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

	log.Fatal(http.ListenAndServe(":8080", router))
}