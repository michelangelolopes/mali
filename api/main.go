package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michelangelolopes/mali/api/user"
)

func main() {
	//Init Router
	r := mux.NewRouter()
	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	// set our port address
	log.Fatal(http.ListenAndServe(":3001", r))
}
