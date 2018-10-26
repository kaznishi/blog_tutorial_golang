package repository

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleRepository interface {
	GetById (id int) *data_model.Article
}
