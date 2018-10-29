package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/kaznishi/blog_tutorial_golang/service"
	"net/http"
	"time"
)

type SessionMiddleware struct {
	sessionService service.SessionService
}

func NewSessionMiddleware(service service.SessionService) SessionMiddleware {
	return SessionMiddleware{sessionService:service}
}

func (m *SessionMiddleware) Run(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := m.getCurrentSession(r)
		if err != nil {
			session, _ = m.sessionService.NewSession(r, "sessId")
			m.sessionService.SaveSession(r, w, session)
		}

		next.ServeHTTP(w, r)
	}
}

func (m *SessionMiddleware) getCurrentSession(r *http.Request) (*sessions.Session, error) {
	cookie, err := r.Cookie("sessId")
	if err != nil {
		return &sessions.Session{}, err
	}

	if cookie.Expires.Before(time.Now()) {
		return &sessions.Session{}, err
	}

	session, err := m.sessionService.GetSession(r, cookie.Name)
	if err != nil {
		return &sessions.Session{}, err
	}

	return session, nil
}
