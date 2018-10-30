package controller

import (
	"fmt"
	"github.com/kaznishi/blog_tutorial_golang/service"
	"net/http"
)

type LoginController struct {
	SessionService service.SessionService
}

func NewLoginController(s service.SessionService) LoginController {
	return LoginController{SessionService: s}
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := initializeTemplate().ParseFiles("view/login.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "ログインページ,エラー")
		fmt.Fprint(w, err)
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "ログインフォーム,エラー")
			fmt.Fprint(w, err)
		}

		err := lc.SessionService.Login(w, r)
		if err != nil {
			fmt.Fprint(w, "ログインフォーム,エラー")
			fmt.Fprint(w, err)
		} else {
			http.Redirect(w, r, "/admin/", 301)
		}
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

func (lc *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	lc.SessionService.Logout(w, r)
	http.Redirect(w, r, "/", 301)
}