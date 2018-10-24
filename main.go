package main

import (
	"github.com/kaznishi/blog_tutorial_golang/config"
	"net/http"
)

func main() {
	router := config.Router()
	http.ListenAndServe(":8080", router)
}
