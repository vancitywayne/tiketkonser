package handlers

import (
    // "html/template"
    "net/http"
    "proyek1-be/models"
    "log"
    "fmt"
    "strconv"
)

func FormatRupiah(price float64) string {
    return fmt.Sprintf("Rp %.2f", price)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    ticketStocks, err := models.GetAllTicketStocks()
    if err != nil {
        http.Error(w, "Unable to fetch ticket stocks", http.StatusInternalServerError)
        return
    }

    data := struct {
        StockRegular int
        StockPremium int
        StockVIP     int
    }{}

    for _, ticketStock := range ticketStocks {
        switch ticketStock.Event {
        case "regular":
            data.StockRegular = ticketStock.Stock
        case "premium":
            data.StockPremium = ticketStock.Stock
        case "vip":
            data.StockVIP = ticketStock.Stock
        }
    }

    if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
    ticketStocks, err := models.GetAllTicketStocks()
    if err != nil {
        log.Printf("Error fetching ticket stocks: %v", err)
        http.Error(w, "Unable to fetch ticket stocks", http.StatusInternalServerError)
        return
    }

    data := struct {
        Stocks []models.TicketStock
    }{
        Stocks: ticketStocks,
    }

    if err := templates.ExecuteTemplate(w, "form.html", data); err != nil {
        log.Printf("Error rendering form template: %v", err)
        http.Error(w, "Unable to render form", http.StatusInternalServerError)
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

    var pricePerTicket float64
    var quantity int
    var err error

    if quantity, err = strconv.Atoi(quantityStr); err != nil || quantity < 1 {
        http.Error(w, "Invalid quantity", http.StatusBadRequest)
        return
    }

    ticketStock, err := models.GetTicketStockByEvent(event)
    if err != nil || ticketStock.Stock < quantity {
        http.Error(w, "Not enough tickets available", http.StatusBadRequest)
        return
    }

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

    if err := models.UpdateTicketStock(event, quantity); err != nil {
        http.Error(w, "Unable to update stock", http.StatusInternalServerError)
        return
    }

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



