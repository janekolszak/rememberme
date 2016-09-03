package rememberme

import (
	"net/http"
	"time"
)

const (
	rememberMeCookieName = "remember"
)

type Rememberme struct {
	Store  Store
	MaxAge time.Duration
}

func (c *Rememberme) Check(r *http.Request) (selector, user string, err error) {
	var now = time.Now()

	l, err := GetLoginCookie(r, rememberMeCookieName)
	if err != nil {
		return
	}

	// TODO: Validate selector, shouldn't be too long etc.

	var hash string
	user, hash, expires, err := c.Store.Get(l.Selector)
	if err != nil {
		return
	}

	if expires.Before(now) {
		err = ErrorSessionExpired
		return
	}

	if !l.Check(hash) {
		err = ErrorBadRequest
	}

	selector = l.Selector

	return
}

func (c *Rememberme) SetCookie(w http.ResponseWriter, r *http.Request, user string) (err error) {
	l := LoginCookie{
		CookieName: rememberMeCookieName,
		MaxAge:     c.MaxAge,
	}

	hash, err := l.GenerateValidator()
	if err != nil {
		return
	}

	// First save to the database
	l.Selector, err = c.Store.Insert(user, hash, time.Now().Add(c.MaxAge))
	if err != nil {
		return
	}

	// Then save to the cookie
	err = l.Save(w, r)
	return
}

func (c *Rememberme) UpdateCookie(w http.ResponseWriter, r *http.Request, selector, user string) (err error) {
	l := LoginCookie{
		Selector:   selector,
		CookieName: rememberMeCookieName,
		MaxAge:     c.MaxAge,
	}

	hash, err := l.GenerateValidator()
	if err != nil {
		return
	}

	// First save to the database
	err = c.Store.Update(selector, user, hash, time.Now().Add(c.MaxAge))
	if err != nil {
		return
	}

	// Then save to the cookie
	err = l.Save(w, r)
	return
}
