package handlers

import (
    "html/template"
    "net/http"
    "proyek1-be/models"
    "log"
    "fmt"
    "strconv"
)

func FormatRupiah(price float64) string {
    return fmt.Sprintf("Rp %.2f", price)
}

// Load templates with custom functions
var templates = template.Must(template.New("").Funcs(template.FuncMap{
    "FormatRupiah": FormatRupiah,
}).ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func BookTicketHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    name := r.FormValue("name")
    email := r.FormValue("email")
    event := r.FormValue("event")
    quantityStr := r.FormValue("quantity")

    log.Println("Received event:", event)

    var pricePerTicket float64
    var quantity int
    var err error

    // Convert quantity to int
    if quantity, err = strconv.Atoi(quantityStr); err != nil || quantity < 1 {
        http.Error(w, "Invalid quantity", http.StatusBadRequest)
        return
    }

    // Set harga otomatis berdasarkan event
    switch event {
    case "regular":
        pricePerTicket = 1000000.0 // 1 juta
    case "premium":
        pricePerTicket = 1500000.0 // 1,5 juta
    case "vip":
        pricePerTicket = 2500000.0 // 2,5 juta
    default:
        http.Error(w, "Invalid event", http.StatusBadRequest)
        return
    }

    totalPrice := pricePerTicket * float64(quantity)

    ticket := models.Ticket{
        Name:      name,
        Email:     email,
        Event:     event,
        Quantity:  quantity,
        TotalPrice: totalPrice,
    }

    if err := ticket.Save(); err != nil {
        http.Error(w, "Unable to book ticket", http.StatusInternalServerError)
        return
    }

    if err := templates.ExecuteTemplate(w, "success.html", ticket); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
