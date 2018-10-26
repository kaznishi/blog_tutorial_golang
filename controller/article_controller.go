package controller

import (
	"fmt"
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
	fmt.Fprint(w, "一覧ページ\n")
	for _, article := range articles {
		fmt.Fprint(w, article.ID)
		fmt.Fprint(w, ", ")
		fmt.Fprint(w, article.Title)
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, article.Content)
		fmt.Fprint(w, "\n======\n")
	}
}
