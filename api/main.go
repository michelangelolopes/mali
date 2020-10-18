package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// set our port address
	log.Fatal(http.ListenAndServe(":3001", r))
}
