package handlers

import (
    "html/template"
    "net/http"
    "proyek1-be/models"
    "log"
	"strconv"
	"github.com/gorilla/mux"
)


type DashboardData struct {
    Tickets     []models.Ticket
    TicketStocks []models.TicketStock
}

var adminTemplates = template.Must(template.New("").Funcs(template.FuncMap{
    "FormatRupiah": FormatRupiah,
}).ParseGlob("templates/*.html"))

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
    tickets, err := models.GetAllTickets()
    if err != nil {
        log.Printf("Error fetching tickets: %v", err)
        http.Error(w, "Unable to fetch tickets", http.StatusInternalServerError)
        return
    }

    ticketStocks, err := models.GetAllTicketStocks()
    if err != nil {
        log.Printf("Error fetching ticket stocks: %v", err)
        http.Error(w, "Unable to fetch ticket stocks", http.StatusInternalServerError)
        return
    }

    data := DashboardData{
        Tickets:     tickets,
        TicketStocks: ticketStocks,
    }

    log.Printf("Dashboard data: %+v", data) // Logging the data to check it

    if err := templates.ExecuteTemplate(w, "dashboard.html", data); err != nil {
        log.Printf("Error rendering template: %v", err)
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
    }
}

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        email := r.FormValue("email")
        event := r.FormValue("event")
        quantity, _ := strconv.Atoi(r.FormValue("quantity"))
        totalPrice, _ := strconv.ParseFloat(r.FormValue("total_price"), 64)

        ticket := models.Ticket{Name: name, Email: email, Event: event, Quantity: quantity, TotalPrice: totalPrice}
        if err := ticket.Save(); err != nil {
            log.Printf("Error creating ticket: %v", err)
            http.Error(w, "Unable to create ticket", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
    } else {
        templates.ExecuteTemplate(w, "create.html", nil)
    }
}

func UpdateTicketHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
        return
    }

    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        email := r.FormValue("email")
        event := r.FormValue("event")
        quantity, _ := strconv.Atoi(r.FormValue("quantity"))

        var pricePerTicket float64

        switch event {
        case "regular":
            pricePerTicket = 1000000.0
        case "premium":
            pricePerTicket = 1500000.0
        case "vip":
            pricePerTicket = 2500000.0
        default:
            http.Error(w, "Invalid event", http.StatusBadRequest)
            return
        }

        totalPrice := pricePerTicket * float64(quantity)

        ticket := models.Ticket{ID: id, Name: name, Email: email, Event: event, Quantity: quantity, TotalPrice: totalPrice}
        if err := ticket.Update(); err != nil {
            log.Printf("Error updating ticket: %v", err)
            http.Error(w, "Unable to update ticket", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
    } else {
        ticket, err := models.GetTicketByID(id)
        if err != nil {
            log.Printf("Error fetching ticket: %v", err)
            http.Error(w, "Unable to fetch ticket", http.StatusInternalServerError)
            return
        }
        templates.ExecuteTemplate(w, "update.html", ticket)
    }
}

func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
        return
    }

    if err := models.DeleteTicket(id); err != nil {
        log.Printf("Error deleting ticket: %v", err)
        http.Error(w, "Unable to delete ticket", http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
