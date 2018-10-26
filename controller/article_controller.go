package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kaznishi/blog_tutorial_golang/service"
)

type ArticleController struct {
	ArticleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return ArticleController{
		ArticleService: articleService,
	}
}

func (c *ArticleController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := c.ArticleService.GetList()
	if err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}

	tmpl, err := template.ParseFiles("view/layout.html.tmpl", "view/article/index.html.tmpl")
	if err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}

	data := struct {
		Title string
		Articles []*data_model.Article
	}{
		Title: "一覧ページ",
		Articles: articles,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Fprint(w, "一覧ページ,エラー")
		fmt.Fprint(w, err)
	}
}

func (c *ArticleController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Errorf("Error : %s", err)
		fmt.Fprint(w, "Error : ", err)
		return
	}
	article, err := c.ArticleService.GetById(id)
	if err != nil {
		fmt.Errorf("Error : %s", err)
		fmt.Fprint(w, "Error : ", err)
		return
	}

	tmpl, err := template.ParseFiles("view/layout.html.tmpl", "view/article/view.html.tmpl")
	if err != nil {
		fmt.Errorf("Error : %s", err)
		fmt.Fprint(w, "Error : ", err)
		return
	}

	data := struct {
		Title string
		Article *data_model.Article
	}{
		Title: "個別ページ",
		Article: article,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Errorf("Error : %s", err)
		fmt.Fprint(w, "Error : ", err)
		return
	}
}