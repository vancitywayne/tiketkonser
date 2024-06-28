package routes

import (
    "github.com/gorilla/mux"
    "proyek1-be/handlers"
    // "net/http"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.LoginHandler).Methods("GET", "POST")
    r.HandleFunc("/signup", handlers.SignupHandler).Methods("GET", "POST")
    r.HandleFunc("/form.html", handlers.FormHandler).Methods("GET")
    r.HandleFunc("/book", handlers.BookTicketHandler).Methods("POST")
    r.HandleFunc("/admin/dashboard", handlers.AdminDashboardHandler).Methods("GET")
    r.HandleFunc("/admin/ticket-stocks", handlers.StockDashboardHandler).Methods("GET")
    r.HandleFunc("/admin/create", handlers.CreateTicketHandler).Methods("GET", "POST")
    r.HandleFunc("/admin/update/{id}", handlers.UpdateTicketHandler).Methods("GET", "POST")
    r.HandleFunc("/admin/delete/{id}", handlers.DeleteTicketHandler).Methods("POST")
    return r
}
