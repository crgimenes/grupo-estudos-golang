package session

import (
	"crypto/rand"
	"net/http"
	"time"
)

const cookieName = "session_example"

type SessionData struct {
	UserID   int
	ExpireAt time.Time
}

type Control struct {
	SessionDataMap map[string]SessionData
}

func New() *Control {
	return &Control{
		SessionDataMap: make(map[string]SessionData),
	}
}

func (c *Control) Get(r *http.Request) (string, *SessionData, bool) {
	cookies := r.Cookies()
	if len(cookies) == 0 {
		return "", nil, false
	}

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", nil, false
	}

	s, ok := c.SessionDataMap[cookie.Value]
	if !ok {
		return "", nil, false
	}

	if s.ExpireAt.Before(time.Now()) {
		delete(c.SessionDataMap, cookie.Value)
		return "", nil, false
	}

	return cookie.Value, &s, true
}

func (c *Control) Delete(w http.ResponseWriter, id string) {
	delete(c.SessionDataMap, id)
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func (c *Control) Save(w http.ResponseWriter, id string, sessionData *SessionData) {
	expireAt := time.Now().Add(3 * time.Hour)
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   id,
		Expires: expireAt,
	}
	sessionData.ExpireAt = expireAt
	c.SessionDataMap[id] = *sessionData

	http.SetCookie(w, cookie)
}

func (c *Control) Create(userID int) (*SessionData, string) {
	sessionData := &SessionData{
		UserID:   userID,
		ExpireAt: time.Now().Add(3 * time.Hour),
	}

	return sessionData, RandomID()
}

func (c *Control) RemoveExpired() {
	for k, v := range c.SessionDataMap {
		if v.ExpireAt.Before(time.Now()) {
			delete(c.SessionDataMap, k)
		}
	}
}

func RandomID() string {
	const (
		length  = 16
		charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	)
	lenCharset := byte(len(charset))
	b := make([]byte, length)
	rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = charset[b[i]%lenCharset]
	}
	return string(b)
}
