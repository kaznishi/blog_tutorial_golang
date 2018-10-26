package controller

import (
	"fmt"
	"github.com/kaznishi/blog_tutorial_golang/service"
	"html/template"
	"net/http"
)

type AdminController struct {
	ArticleService service.ArticleService
}

func NewAdminController(articleService service.ArticleService) AdminController {
	return AdminController{
		ArticleService: articleService,
	}
}

func (c *AdminController) Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("view/layout_admin.html.tmpl", "view/admin/index.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}

	data := struct {
		Title string
	}{
		Title: "管理画面トップページ",
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}
}

