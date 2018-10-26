package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleMySQLDAO struct {
	MySQLConn *sql.DB
}

func (dao *ArticleMySQLDAO) GetList() ([]*data_model.Article, error) {
	result := make([]*data_model.Article, 0)
	a1 := &data_model.Article{ID:1, Title: "hoge", Body: "hogebody"}
	result = append(result, a1)
	a2 := &data_model.Article{ID:2, Title: "fuga", Body: "fugabody"}
	result = append(result, a2)
	a3 := &data_model.Article{ID:3, Title: "hoga", Body: "hogabody"}
	result = append(result, a3)
	return result, nil // dummy
}

func (dao *ArticleMySQLDAO) GetById(id int) (*data_model.Article, error) {
	return &data_model.Article{}, nil // dummy
}
