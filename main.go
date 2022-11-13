package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IRFAN374/goRestApi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/allgrocery", handlers.AddGrocery).Methods("GET")
	router.HandleFunc("/grocery/{name}", handlers.GetGrocery).Methods("GET")
	router.HandleFunc("/grocery", handlers.AddGrocery).Methods("POST")
	router.HandleFunc("/grocery/{name}", handlers.UpdateGrocery).Methods("PUT")
	router.HandleFunc("/grocery/{name}", handlers.DeleteGrocery).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
