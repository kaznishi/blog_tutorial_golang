package middleware

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/kaznishi/blog_tutorial_golang/service"
	"net/http"
)

type SessionMiddleware struct {
	sessionService service.SessionService
}

func NewSessionMiddleware(service service.SessionService) SessionMiddleware {
	return SessionMiddleware{sessionService:service}
}

func (m *SessionMiddleware) SessionStart(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := m.getCurrentSession(r)
		if err != nil {
			session, _ = m.sessionService.NewSession(r, "sessId")
			m.sessionService.SaveSession(r, w, session)
		}

		next.ServeHTTP(w, r)
	}
}

func (m *SessionMiddleware) ForOnlyLoginUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !m.sessionService.IsLogin(r) {
			cookie := &http.Cookie{
				Name: "Pragma",
				Value: "no-cache",
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/login", 301)
		}
		next.ServeHTTP(w, r)
	}
}

func (m *SessionMiddleware) getCurrentSession(r *http.Request) (*sessions.Session, error) {
	cookie, err := r.Cookie("sessId")
	if err != nil {
		return &sessions.Session{}, err
	}

	session, err := m.sessionService.GetSession(r, cookie.Name)
	if err != nil {
		return &sessions.Session{}, err
	}

	fmt.Println(session)
	return session, nil
}
