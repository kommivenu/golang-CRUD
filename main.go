package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/getallusers", GetUsers).Methods("GET")
	r.HandleFunc("/getuser/{id}", GetUser).Methods("GET")
	r.HandleFunc("/saveuser", CreateUser).Methods("POST")
	r.HandleFunc("/updateuser/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/deleteuser/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9090", r))

}

func main() {
	InitialMigration()
	initializeRouter()

}
