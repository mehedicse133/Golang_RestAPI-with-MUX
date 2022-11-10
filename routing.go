package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func hendler() {
	r := mux.NewRouter()

	r.HandleFunc("/students", add_student).Methods("POST")
	r.HandleFunc("/students", get_all_students).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent_by_id).Methods("GET")
	r.HandleFunc("/students/{id}", update_student).Methods("PUT")
	r.HandleFunc("/students/{id}", delete_student).Methods("DELETE")
	http.ListenAndServe(":6001", r)

}
