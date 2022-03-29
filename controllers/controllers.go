package controllers

import (
	"encoding/json"
	"fmt"
	"goact-todo/middlewares"
	"goact-todo/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	todos, err := middlewares.GetAllTasks()
	if err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var task models.Todo

	json.NewDecoder(r.Body).Decode(&task)
	task.ID = primitive.NewObjectID()

	if err := middlewares.AddNewTaskToDB(task); err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(task)
}

func TaskDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("Done:", id)

	if err := middlewares.TaskDone(id); err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("Undone:", id)

	if err := middlewares.TaskUnDone(id); err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("Delete:", id)

	if err := middlewares.DeleteTaskFromDB(id); err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := middlewares.DeleteAllTaskFromDB(); err != nil {
		log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}
