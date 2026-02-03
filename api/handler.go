package api

import (
	"app/internal/task"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
)

/*
func ShowMenu(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
	w.Write([]byte(`
	<h1>Меню<h1>
	<a href="/tasks/create">Создать задачу</a><br>
	<a href="/tasks">Показать задачи</a><br>
	<a href="/exit">Выход</a>
	`))
	}
}*/

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var t task.CreateTaskRequest
	var newTask task.Task

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println("Сработала ошибка", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	text := strings.TrimSpace(t.Text)
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task text cannot be empty"))
		return
	}

	newTask = task.Task{
		ID:   len(task.List) + 1,
		Text: text,
		Done: false,
	}

	task.List = append(task.List, newTask)

	log.Printf("Task created ID=%d, Text=%s", newTask.ID, newTask.Text)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created"))
}

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	err := json.NewEncoder(w).Encode(task.List)
	if err != nil {
		log.Println("сработала ошибка", err)
		return
	}

	log.Println("ShowTasks: sent", len(task.List), "tasks")
	pp.Println(task.List)
}
