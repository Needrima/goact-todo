package main

import (
	"goact-todo/router"
	"net/http"
)

func main() {
	r := router.Router()

	http.ListenAndServe(":8080", r)
}
