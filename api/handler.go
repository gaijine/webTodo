package api

import (
	"app/internal/task"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

var nextID = 1

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var t task.CreateTaskRequest
	var newTask task.Task
	/*
		if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}*/

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
		ID:   nextID,
		Text: text,
		Done: false,
	}

	task.List = append(task.List, newTask)
	nextID++

	log.Printf("Task created ID=%d, Text=%s", newTask.ID, newTask.Text)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created"))
}

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}*/

	err := json.NewEncoder(w).Encode(task.List)
	if err != nil {
		log.Println("сработала ошибка", err)
		return
	}

	log.Println("ShowTasks: sent", len(task.List), "tasks")
	pp.Println(task.List)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	if len(task.List) == 0 {
		log.Println("[WARN] Task list is empty")
		w.Write([]byte("Список задач пуст"))
		return
	}

	var change task.UpdateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&change)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newText := strings.TrimSpace(change.Text)
	if newText == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task text cannot be empty"))
		return
	}
	/*
		if change.ID <= 0 || change.ID > len(task.List) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
			return
		}*/

	// Найти индекс по уникальному ID
	idx := -1
	for i, task := range task.List {
		if task.ID == change.ID {
			idx = i
			break
		}
	}

	if idx == -1 {
		http.Error(w, "task not found", http.StatusBadRequest)
		return
	}

	oldText := task.List[idx].Text
	task.List[idx].Text = newText

	log.Printf("[INFO] Task <%s> changed to <%s> successfully", oldText, newText)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task update successfully"))

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	if len(task.List) == 0 {
		log.Println("[WARN] Task list is empty")
		w.Write([]byte("Список задач пуст"))
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("task ID don`t be empty"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid task id"))
		return
	}
	/*
		if id <= 0 || id > len(task.List) {
			http.Error(w, "такой задачи нет", http.StatusBadRequest)
			return
		}*/

	idx := -1
	for i, task := range task.List {
		if task.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		http.Error(w, "Такой задачи нет", http.StatusBadRequest)
		return
	}

	deletedTask := task.List[idx].Text
	task.List = append(task.List[:idx], task.List[idx+1:]...)
	log.Printf("[INFO] task deleted: id=%d, text=<%s>", id, deletedTask)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}

//////////////////////////////////////////////////////////
/*
type Users struct{
	Login string
	Password string
	Email string
}
type userEmail struct{
	Login string
	Email string
}

var usr = make(map[string]Users)
func main(){
r := mux.NewRouter()

r.HandleFunc("/user", getUsers).Methods("GET")
r.HandleFunc("/user", createUser).Methods("POST").Headers("Content-Type", "application/json")
r.HandleFunc("/user", updateUser).Methods("PUT").Headers("Content-Type", "application/json")
r.HandleFunc("/user", deleteUser).Methods("DELETE")

http.ListenAndServe(":8080", r)
}

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func createUser(w http.ResponseWriter, r *http.Request){
var u Users
json.NewDecoder(r.Body).Decode(&u)
usr[u.Login] = u

w.Write([]byte("пользователь создан"))
}

func updateUser(w http.ResponseWriter, r *http.Request){
var u userEmail
json.NewDecoder(r.Body).Decode(&u)

v, ok := usr[u.Login]
if !ok{
	http.Error(w, "такого пользователя нет", http.StatusBadRequest)
	return
}
v.Email = u.Email
usr[v.Login] = v

}

func deleteUser(w http.ResponseWriter, r *http.Request){
	login := r.URL.Query().Get("login")
	if login == ""{
		http.Error(w, "такого пользователя нет", http.StatusBadRequest)
	return
	}
	delete(usr, login)
}

func middleware(next http.Handler) http.HandlerFunc{
 return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
 })
}
*/
