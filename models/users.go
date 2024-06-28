package models

import (
    "proyek1-be/database"
    "golang.org/x/crypto/bcrypt"
    "log"
)

type User struct {
    Username string
    Password string
}

func GetUserByUsername(username string) (*User, error) {
    db := database.GetDB()
    row := db.QueryRow("SELECT username, password FROM users WHERE username=?", username)

    var user User
    if err := row.Scan(&user.Username, &user.Password); err != nil {
        log.Printf("Error fetching user from database: %v", err)
        return nil, err
    }

    return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (u *User) Save() error {
    db := database.GetDB()
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
    if err != nil {
        log.Printf("Error saving user to database: %v", err)
        return err
    }
    return nil
}
