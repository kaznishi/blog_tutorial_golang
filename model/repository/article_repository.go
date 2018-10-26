package repository

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleRepository interface {
	GetList() ([]*data_model.Article, error)
	GetById(id int) (*data_model.Article, error)
}
