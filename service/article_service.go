package service

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
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

func (as *ArticleService) GetList() ([]*data_model.Article, error) {
	return as.ArticleRepository.GetList()
}

func (as *ArticleService) GetById(id int) (*data_model.Article, error) {
	return as.ArticleRepository.GetById(id)
}

func (as *ArticleService) CreateArticle(a *data_model.Article) (int, error) {
	return as.ArticleRepository.Create(a)
}