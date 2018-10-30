package service

import (
	"github.com/gorilla/sessions"
	"github.com/kaznishi/blog_tutorial_golang/errors"
	"github.com/kaznishi/blog_tutorial_golang/model/repository"
	"github.com/kaznishi/blog_tutorial_golang/util"
	"net/http"
)

type SessionService struct {
	sessionStore sessions.Store
	userRepository repository.UserRepository
}

func NewSessionService(store sessions.Store, userRepository repository.UserRepository) SessionService {
	return SessionService{
		sessionStore: store,
		userRepository: userRepository,
	}
}

func (s *SessionService) NewSession(r *http.Request, name string) (*sessions.Session, error) {
	return s.sessionStore.New(r, name)
}

func (s *SessionService) GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return s.sessionStore.Get(r, name)
}

func (s *SessionService) SaveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return s.sessionStore.Save(r, w, session)
}

func (s *SessionService) Login(name string, password string, w http.ResponseWriter, r *http.Request) error {
	// ユーザ照合(無効ならreturn error)
	user, err := s.userRepository.GetByName(name)
	if err != nil {
		return err
	}
	if user.Password != util.PasswordHashing(password, user.Salt) {
		return errors.NOT_FOUND_ERROR
	}

	// セッションにログイン情報を書き込み
	session, err := s.GetSession(r, "sessId")
	if err != nil {
		return err
	}
	session.Values["login"] = "true"
	session.Values["user"] = user.ID

	return s.SaveSession(r, w, session)
}

func (s *SessionService) Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := s.GetSession(r, "sessId")
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1

	return s.SaveSession(r, w, session)
}

func (s *SessionService) IsLogin(r *http.Request) bool {
	session, err := s.GetSession(r, "sessId")
	if err != nil {
		return false
	}
	if session.Values["login"] != "true" {
		return false
	}
	return true
}