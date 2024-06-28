package handlers

import (
    "net/http"
    "proyek1-be/models"
    "github.com/gorilla/sessions"
    "golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        user, err := models.GetUserByUsername(username)
        if err != nil || !models.CheckPasswordHash(password, user.Password) {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        session, _ := store.Get(r, "session-name")
        session.Values["authenticated"] = true
        session.Save(r, w)

        http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
    } else {
        http.ServeFile(w, r, "templates/index.html")
    }
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("new_username")
        password := r.FormValue("new_password")

        if _, err := models.GetUserByUsername(username); err == nil {
            http.Error(w, "Username already taken", http.StatusBadRequest)
            return
        }

        hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(w, "Unable to create account", http.StatusInternalServerError)
            return
        }

        user := models.User{
            Username: username,
            Password: string(hash),
        }

        if err := user.Save(); err != nil {
            http.Error(w, "Unable to create account", http.StatusInternalServerError)
            return
        }

        session, _ := store.Get(r, "session-name")
        session.Values["authenticated"] = true
        session.Save(r, w)

        http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
    } else {
        http.ServeFile(w, r, "templates/signup.html")
    }
}
