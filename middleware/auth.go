package middleware

import (
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func RequireLogin(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "session-name")
        if session.Values["authenticated"] != true {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
    })
}
