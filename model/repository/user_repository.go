package repository

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
)

type UserRepository interface {
	GetList() ([]*data_model.User, error)
	GetById(id int) (*data_model.User, error)
	GetByName(name string) (*data_model.User, error)
	Create(*data_model.User) (int, error)
}
