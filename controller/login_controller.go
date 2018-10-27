package controller

import (
	"fmt"
	"net/http"
)

type LoginController struct {
}

func NewLoginController() LoginController {
	return LoginController{
	}
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := initializeTemplate().ParseFiles("view/login.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "ログインページ,エラー")
		fmt.Fprint(w, err)
	}

	data := struct {
		Title string
	}{
		Title: "ログインフォーム",
	}

	if err := tmpl.ExecuteTemplate(w, "login.html.tmpl", data); err != nil {
		fmt.Fprint(w, "ログインページ,エラー")
		fmt.Fprint(w, err)
	}

}
