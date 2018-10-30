package infrastructure

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/errors"
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type UserMySQLDAO struct {
	MySQLConn *sql.DB
}

func (dao *UserMySQLDAO) fetch(query string, args ...interface{}) ([]*data_model.User, error) {
	rows, err := dao.MySQLConn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*data_model.User, 0)
	for rows.Next() {
		a := new(data_model.User)
		err = rows.Scan(
			&a.ID,
			&a.Name,
			&a.Password,
			&a.Salt,
			&a.IsActive,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (dao *UserMySQLDAO) GetList() ([]*data_model.User, error) {
	query := `SELECT id, name, password, salt, is_active FROM users`
	return dao.fetch(query)
}

func (dao *UserMySQLDAO) GetById(id int) (*data_model.User, error) {
	query := `SELECT id, name, password, salt, is_active FROM users WHERE id = ?`
	list, err := dao.fetch(query, id)
	if err != nil {
		return nil, err
	}

	result := &data_model.User{}
	if len(list) > 0 {
		result = list[0]
	} else {
		return nil, errors.NOT_FOUND_ERROR
	}

	return result, nil
}

func (dao *UserMySQLDAO) GetByName(name string) (*data_model.User, error) {
	query := `SELECT id, name, password, salt, is_active FROM users WHERE name = ?`
	list, err := dao.fetch(query, name)
	if err != nil {
		return nil, err
	}

	result := &data_model.User{}
	if len(list) > 0 {
		result = list[0]
	} else {
		return nil, errors.NOT_FOUND_ERROR
	}

	return result, nil
}

func (dao *UserMySQLDAO) Create(a *data_model.User) (int, error) {
	query := `INSERT users SET name = ?, password = ?, salt = ?, is_active = ?`
	stmt, err := dao.MySQLConn.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(
		a.Name,
		a.Password,
		a.Salt,
		a.IsActive)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), nil
}
