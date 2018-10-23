package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "一覧ページ")
}

func Article(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "記事個別ページ ID:")
	fmt.Fprint(w, vars["id"])
}
