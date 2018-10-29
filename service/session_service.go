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

func (s *SessionService) GetSession(r *http.Request, name string) (*sessions.Session, error){
	return s.sessionStore.Get(r, name)
}

func (s *SessionService) SaveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return s.sessionStore.Save(r, w, session)
}

