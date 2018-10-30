package service

import (
	"github.com/kaznishi/blog_tutorial_golang/model/data_model"
	"github.com/kaznishi/blog_tutorial_golang/model/repository"
	"github.com/kaznishi/blog_tutorial_golang/util"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetList() ([]*data_model.User, error) {
	return us.UserRepository.GetList()
}

func (us *UserService) CreateUser(name string, password string) (int, error) {
	// salt生成、passwordハッシュ化、is_activeに1を固定で差込
	user := new(data_model.User)
	user.Name = name
	user.Salt = util.RandomString()
	user.Password = util.PasswordHashing(password, user.Salt)
	user.IsActive = 1

	return us.UserRepository.Create(user)
}

