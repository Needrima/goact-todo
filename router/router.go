package router

import (
	"campmart/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middlewares.GetAllTasks).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/task", middlewares.CreateNewTask).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/task/done/{id}", middlewares.TaskDone).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/task/undo/{id}", middlewares.UndoTask).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/task/delete/{id}", middlewares.DeleteTask).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/deletealltasks", middlewares.DeleteAllTasks).Methods(http.MethodDelete, http.MethodOptions)

	return r
}
