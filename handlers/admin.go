package handlers

import (
    "html/template"
    "net/http"
    "proyek1-be/models"
    "log"
	"strconv"
	"github.com/gorilla/mux"
)

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
    tickets, err := models.GetAllTickets()
    if err != nil {
        log.Printf("Error fetching tickets: %v", err)
        http.Error(w, "Unable to fetch tickets", http.StatusInternalServerError)
        return
    }

    templates := template.Must(template.ParseFiles("templates/dashboard.html"))
    if err := templates.ExecuteTemplate(w, "dashboard.html", tickets); err != nil {
        log.Printf("Error rendering template: %v", err)
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
    }
}

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        email := r.FormValue("email")
        event := r.FormValue("event")

        ticket := models.Ticket{Name: name, Email: email, Event: event}
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

        ticket := models.Ticket{ID: id, Name: name, Email: email, Event: event}
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
