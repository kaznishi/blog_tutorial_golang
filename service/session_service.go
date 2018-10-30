package service

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type SessionService struct {
	sessionStore sessions.Store
}

func NewSessionService(store sessions.Store) SessionService {
	return SessionService{sessionStore: store}
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

func (s *SessionService) Login(w http.ResponseWriter, r *http.Request) error {
	// ユーザ照合(無効ならreturn error)

	// セッションにログイン情報を書き込み
	session, err := s.GetSession(r, "sessId")
	if err != nil {
		return err
	}
	session.Values["login"] = "true"
	session.Values["user"] = "1"

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
