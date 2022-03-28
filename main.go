package main

import (
	"fmt"
	"goact-todo/router"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Running on port :8080, visit localhost:8080...")
	http.ListenAndServe(":8080", r)
}
