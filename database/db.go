package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
    var err error
    // connStr := "user=username dbname=proyek1db sslmode=disable"
    dsn := "root:@tcp(127.0.0.1:3306)/proyek1db"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}

func GetDB() *sql.DB {
    return db
}
