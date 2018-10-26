package service

import (
	"github.com/kaznishi/blog_tutorial_golang/model/repository"
)

type ArticleService struct {
	ArticleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) ArticleService {
	return ArticleService{
		ArticleRepository: articleRepository,
	}
}
