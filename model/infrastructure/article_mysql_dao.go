package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/errors"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
	"time"
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

func (dao *ArticleMySQLDAO) Create(a *data_model.Article) (int, error) {
	query := `INSERT articles SET title = ?, content = ?, created_at = ?, updated_at = ?`
	stmt, err := dao.MySQLConn.Prepare(query)
	if err != nil {
		return 0, err
	}
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	res, err := stmt.Exec(
		a.Title,
		a.Content,
		a.CreatedAt.Format("2006/01/02 15:04:05"),
		a.UpdatedAt.Format("2006/01/02 15:04:05"))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), nil
}