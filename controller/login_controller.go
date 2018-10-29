package controller

import (
	"fmt"
	"net/http"
)

//var cStore = sessions.NewCookieStore([]byte("a6b0e040989e6131daccca9290cb64a0444b52dfc3bf22b8b77f938542f79757"))

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

	//session, _ := cStore.Get(r, "session-name")
	//session.Values["foo"] = "bar"
	//session.Values[42] = 43
	//session.Save(r, w)

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
