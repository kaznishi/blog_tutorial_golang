package controller

import (
	"fmt"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
	"html/template"
	"net/http"

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
