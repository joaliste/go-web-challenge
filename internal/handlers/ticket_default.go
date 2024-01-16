package handlers

import (
	"app/internal"
	"app/platform/web/response"
	"errors"
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
			switch {
			case errors.Is(err, internal.ErrCountryNotFound):
				response.Text(w, http.StatusBadRequest, "country does not exist")
			default:
				response.Text(w, http.StatusBadGateway, "internal error")
			}

			return
		}

		// - response
		response.JSON(w, http.StatusOK, ticketAmount)
	}
}

func (d *DefaultTickets) GetTicketAverageByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		country := chi.URLParam(r, "destination")

		// process
		percentage, err := d.sv.GetTicketAverageByDestinationCountry(country)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrCountryNotFound):
				response.Text(w, http.StatusBadRequest, "country does not exist")
			default:
				response.Text(w, http.StatusBadGateway, "internal error")
			}

			return
		}

		// - response
		response.JSON(w, http.StatusOK, percentage)
	}
}
