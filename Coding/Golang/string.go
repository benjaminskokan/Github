package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Person struct {
    ID string json:"id, omitempty"
    Firstname string json:"firstname, omitempty"
    Lastname string json:"lastname, omitempty"
}


func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
    router:= mux.NewRouter()
    router.HandleFunc("/people", GetPeopleEndpoint.Methods("GET"))
    router.HandleFunc("/people/{id}", GetPersonEndpoint.Methods("GET"))
    router.HandleFunc("/people/{id}", CreatePeopleEndpoint.Methods("POST"))
    router.HandleFunc("/people", DeletePeopleEndpoint.Methods("DELETE"))
    log.Fatal(http.ListenAndServe("12345", router))
}
