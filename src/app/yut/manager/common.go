package manager

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)),
	[]byte(securecookie.GenerateRandomKey(32)),
	[]byte(securecookie.GenerateRandomKey(32)),
	[]byte(securecookie.GenerateRandomKey(32)),
	)

func init() {
	Store.MaxAge(24 * 3600)
}

func GetSession(name string, r *http.Request) interface{} {
	session, _ := Store.Get(r, "session-key")

	if nameValue, ok := session.Values[name]; ok {
		return nameValue
	}
	return ""
}
