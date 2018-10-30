package repository

import (
	"database/sql"
	"github.com/kaznishi/blog_tutorial_golang/model/infrastructure"
)

type RepositoryManager struct {
	MySQLConn *sql.DB
}

func NewRepositoryManager(mysqlConn *sql.DB) RepositoryManager {
	return RepositoryManager{
		MySQLConn: mysqlConn,
	}
}

func (rm *RepositoryManager) NewArticleRepository() ArticleRepository {
	return &infrastructure.ArticleMySQLDAO{rm.MySQLConn}
}

func (rm *RepositoryManager) NewUserRepository() UserRepository {
	return &infrastructure.UserMySQLDAO{rm.MySQLConn}
}
