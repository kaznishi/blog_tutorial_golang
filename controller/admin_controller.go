package controller

import (
	"fmt"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
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
	articles, err := c.ArticleService.GetList()
	if err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}

	tmpl, err := template.ParseFiles("view/layout_admin.html.tmpl", "view/admin/index.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}

	data := struct {
		Title string
		Articles []*data_model.Article
	}{
		Title: "管理画面トップページ",
		Articles: articles,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}
}

func (c *AdminController) NewArticle(w http.ResponseWriter, r *http.Request) {
	a := new(data_model.Article)

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "新規記事投稿フォーム,エラー")
			fmt.Fprint(w, err)
		}
		a.Title = r.FormValue("title")
		a.Content = r.FormValue("content")

		_, err := c.ArticleService.CreateArticle(a)
		if err != nil {
			fmt.Fprint(w, "新規記事投稿フォーム,エラー")
			fmt.Fprint(w, err)
		} else {
			http.Redirect(w, r, "/admin/", 301)
		}
	}

	tmpl, err := template.ParseFiles("view/layout_admin.html.tmpl", "view/admin/new_article.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "新規記事投稿フォーム,エラー")
		fmt.Fprint(w, err)
	}

	data := struct {
		Title string
		Article *data_model.Article
	}{
		Title: "新規記事投稿フォーム",
		Article: a,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprint(w, "新規記事投稿フォーム,エラー")
		fmt.Fprint(w, err)
	}

}
