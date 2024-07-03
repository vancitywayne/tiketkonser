package models

import (
    "log"
    "proyek1-be/database"
)

type Ticket struct {
    ID         int
    Name       string
    Email      string
    Event      string
    Quantity   int
    TotalPrice float64
}

func (t *Ticket) Save() error {
    db := database.GetDB()
    _, err := db.Exec("INSERT INTO tickets (name, email, event, quantity, total_price) VALUES (?, ?, ?, ?, ?)", t.Name, t.Email, t.Event, t.Quantity, t.TotalPrice)
    return err
}

func (t *Ticket) Update() error {
    db := database.GetDB()
    _, err := db.Exec("UPDATE tickets SET name=?, email=?, event=?, quantity=?, total_price=? WHERE id=?", t.Name, t.Email, t.Event, t.Quantity, t.TotalPrice, t.ID)
    if err != nil {
        log.Printf("Error updating data in database: %v", err)
    }
    return err
}

func GetTicketByID(id int) (*Ticket, error) {
    db := database.GetDB()
    row := db.QueryRow("SELECT id, name, email, event, quantity, total_price FROM tickets WHERE id=?", id)

    var ticket Ticket
    if err := row.Scan(&ticket.ID, &ticket.Name, &ticket.Email, &ticket.Event, &ticket.Quantity, &ticket.TotalPrice); err != nil {
        log.Printf("Error fetching ticket from database: %v", err)
        return nil, err
    }

    return &ticket, nil
}

func GetAllTickets() ([]Ticket, error) {
    db := database.GetDB()
    rows, err := db.Query("SELECT id, name, email, event, quantity, total_price FROM tickets")
    if err != nil {
        log.Printf("Error fetching tickets from database: %v", err)
        return nil, err
    }
    defer rows.Close()

    var tickets []Ticket
    for rows.Next() {
        var ticket Ticket
        if err := rows.Scan(&ticket.ID, &ticket.Name, &ticket.Email, &ticket.Event, &ticket.Quantity, &ticket.TotalPrice); err != nil {
            log.Printf("Error scanning row: %v", err)
            return nil, err
        }
        tickets = append(tickets, ticket)
    }

    return tickets, nil
}

func DeleteTicket(ticketID string) error {
    db := database.GetDB()
    _, err := db.Exec("DELETE FROM tickets WHERE id = ?", ticketID)
    if err != nil {
        log.Printf("Error deleting ticket from database: %v", err)
        return err
    }
    return nil
}

// func GetTicketStock() (int, error) {
//     db := database.GetDB()
//     var stock int
//     err := db.QueryRow("SELECT SUM(quantity) FROM tickets").Scan(&stock)
//     if err != nil {
//         log.Printf("Error fetching stock from database: %v", err)
//         return 0, err
//     }
//     return stock, nil
// }
