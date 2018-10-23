package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ログイン")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ログアウト")
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "管理画面トップ")
}

func AdminListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "管理画面ユーザ一覧")
}

func AdminNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "管理画面新規ユーザ")
}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "管理画面ユーザ削除")
	fmt.Fprint(w, vars["id"])
}

func AdminViewArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "管理画面記事閲覧")
	fmt.Fprint(w, vars["id"])
}

func AdminEditArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "管理画面記事編集Get")
	fmt.Fprint(w, vars["id"])
}

func AdminDeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "管理画面記事削除")
	fmt.Fprint(w, vars["id"])
}
