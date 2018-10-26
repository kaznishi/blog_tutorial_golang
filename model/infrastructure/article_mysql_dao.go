package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleMySQLDAO struct {
	MySQLConn *sql.DB
}

func (dao *ArticleMySQLDAO) GetById (id int) *data_model.Article {
	return &data_model.Article{} // dummy
}
