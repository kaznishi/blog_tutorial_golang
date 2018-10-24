package config

import (
	"github.com/gorilla/mux"
	"github.com/kaznishi/blog_tutorial_golang/handlers"
)

func Router() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", handlers.Index).Methods("GET")
	m.HandleFunc("/article/{id:[0-9]+}", handlers.Article).Methods("GET")
	m.HandleFunc("/login", handlers.Login)
	m.HandleFunc("/logout", handlers.Logout)
	m.HandleFunc("/admin/", handlers.AdminIndex).Methods("GET")
	m.HandleFunc("/admin/user/list", handlers.AdminListUser).Methods("GET")
	m.HandleFunc("/admin/user/new", handlers.AdminNewUser).Methods("POST")
	m.HandleFunc("/admin/user/delete/{id:[0-9]+}", handlers.AdminDeleteUser).Methods("POST")
	m.HandleFunc("/admin/article/view/{id:[0-9]+}", handlers.AdminViewArticle).Methods("GET")
	m.HandleFunc("/admin/article/edit/{id:[0-9]+}", handlers.AdminEditArticle)
	m.HandleFunc("/admin/article/delete/{id:[0-9]+}", handlers.AdminDeleteArticle).Methods("POST")

	return m
}
