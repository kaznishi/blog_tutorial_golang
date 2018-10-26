package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/errors"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type ArticleMySQLDAO struct {
	MySQLConn *sql.DB
}

func (dao *ArticleMySQLDAO) fetch(query string, args ...interface{}) ([]*data_model.Article, error) {
	rows, err := dao.MySQLConn.Query(query, args...)
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
			&a.CreatedAt,
			&a.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (dao *ArticleMySQLDAO) GetList() ([]*data_model.Article, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM articles`
	return dao.fetch(query)
}

func (dao *ArticleMySQLDAO) GetById(id int) (*data_model.Article, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM articles WHERE id = ?`
	list, err := dao.fetch(query, id)
	if err != nil {
		return nil, err
	}

	result := &data_model.Article{}
	if len(list) > 0 {
		result = list[0]
	} else {
		return nil, errors.NOT_FOUND_ERROR
	}

	return result, nil
}
