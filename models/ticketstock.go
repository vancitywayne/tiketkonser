package models

import (
    "proyek1-be/database"
    "log"
    "fmt"
)

type TicketStock struct {
    Event string
    Stock int
}

func GetTicketStockByEvent(event string) (*TicketStock, error) {
    db := database.GetDB()
    row := db.QueryRow("SELECT event, stock FROM ticket_stocks WHERE event=?", event)

    var ticketStock TicketStock
    if err := row.Scan(&ticketStock.Event, &ticketStock.Stock); err != nil {
        log.Printf("Error fetching ticket stock from database: %v", err)
        return nil, err
    }

    return &ticketStock, nil
}

func UpdateTicketStock(event string, quantity int) error {
    db := database.GetDB()
    result, err := db.Exec("UPDATE ticket_stocks SET stock = stock - ? WHERE event = ? AND stock >= ?", quantity, event, quantity)
    if err != nil {
        log.Printf("Error updating ticket stock in database: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("No rows updated, possibly due to insufficient stock")
        return fmt.Errorf("insufficient stock for event: %s", event)
    }

    log.Printf("Stock updated successfully for event: %s, quantity: %d", event, quantity)
    return nil
}

func GetAllTicketStocks() ([]TicketStock, error) {
    db := database.GetDB()
    rows, err := db.Query("SELECT event, stock FROM ticket_stocks")
    if err != nil {
        log.Printf("Error fetching ticket stocks from database: %v", err)
        return nil, err
    }
    defer rows.Close()

    var ticketStocks []TicketStock
    for rows.Next() {
        var ticketStock TicketStock
        if err := rows.Scan(&ticketStock.Event, &ticketStock.Stock); err != nil {
            log.Printf("Error scanning row: %v", err)
            return nil, err
        }
        ticketStocks = append(ticketStocks, ticketStock)
    }

    return ticketStocks, nil
}
