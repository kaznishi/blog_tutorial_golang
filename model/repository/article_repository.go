package repository

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleRepository interface {
	GetList() ([]*data_model.Article, error)
	GetById(id int) (*data_model.Article, error)
	Create(*data_model.Article) (int, error)
	Update(*data_model.Article) (*data_model.Article, error)
}
