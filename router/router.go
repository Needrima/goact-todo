package router

import (
	"goact-todo/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.GetAllTasks).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/task", controllers.CreateNewTask).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/task/done/{id}", controllers.TaskDone).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/task/undo/{id}", controllers.UndoTask).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/task/delete/{id}", controllers.DeleteTask).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/deletealltasks", controllers.DeleteAllTasks).Methods(http.MethodDelete, http.MethodOptions)

	return r
}
