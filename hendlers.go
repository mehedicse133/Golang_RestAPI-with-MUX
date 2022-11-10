package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var students []Student

func add_student(w http.ResponseWriter, r *http.Request) {
	var newstuent Student
	json.NewDecoder(r.Body).Decode(&newstuent)
	students = append(students, newstuent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newstuent)
}
func get_all_students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func getStudent_by_id(w http.ResponseWriter, r *http.Request) {
	var s_id string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	student := students[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func update_student(w http.ResponseWriter, r *http.Request) {

	var stu_id string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(stu_id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	var update_stu Student
	json.NewDecoder(r.Body).Decode(&update_stu)
	students[id] = update_stu

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(update_stu)
}

func delete_student(w http.ResponseWriter, r *http.Request) {
	var stu_id string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(stu_id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	students = append(students[:id], students[id+1:]...)
	w.WriteHeader(200)
}
