package main

import (
    "log"
    "net/http"
    "proyek1-be/routes"
)

func main() {
    r := routes.RegisterRoutes()
    log.Println("Server started on: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
    
}
