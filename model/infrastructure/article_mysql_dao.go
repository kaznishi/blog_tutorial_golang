package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleMySQLDAO struct {
	MySQLConn *sql.DB
}

func (dao *ArticleMySQLDAO) GetList() ([]*data_model.Article, error) {
	query := `SELECT id, title, content FROM articles`

	rows, err := dao.MySQLConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*data_model.Article, 0)
	for rows.Next() {
		a := new(data_model.Article)
		err = rows.Scan(
			&a.ID,
			&a.Title,
			&a.Content,
			)

		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (dao *ArticleMySQLDAO) GetById(id int) (*data_model.Article, error) {
	return &data_model.Article{}, nil // dummy
}
