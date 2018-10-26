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
	fmt.Fprint(w, "一覧ページ")
}
