package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouter() {
	r := mux.NewRouter()
	//r.HandleFunc("/", ShowMenu).Methods("GET")
	r.HandleFunc("/tasks/create", CreateTask).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/tasks", ShowTasks).Methods("GET")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
