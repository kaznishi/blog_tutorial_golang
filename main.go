package main

import (
	"github.com/kaznishi/blog_tutorial_golang/web"
	"net/http"
)

func main() {
	router := web.Router()
	http.ListenAndServe(":8080", router)
}
