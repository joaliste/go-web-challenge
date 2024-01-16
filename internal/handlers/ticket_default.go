package handlers

import (
	"app/internal"
	"app/platform/web/response"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DefaultTickets struct {
	sv internal.TicketService
}

func NewDefaultTickets(sv internal.TicketService) *DefaultTickets {
	return &DefaultTickets{
		sv: sv,
	}
}

func (d *DefaultTickets) GetTicketByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		country := chi.URLParam(r, "destination")

		// process
		ticketAmount, err := d.sv.GetTicketByDestinationCountry(country)

		if err != nil {
			response.Text(w, http.StatusBadRequest, "country does not exist")
		}

		// - response
		response.JSON(w, http.StatusOK, ticketAmount)
	}
}
