package service

import (
	"app/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.TicketRepository
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.TicketRepository) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	return
}

func (s *ServiceTicketDefault) GetTicketByDestinationCountry(country string) (int, error) {
	data := s.rp.GetTicketByDestinationCountry(country)
	var err error = nil
	if len(data) == 0 {
		err = internal.ErrCountryNotFound
	}

	return len(data), err
}

func (s *ServiceTicketDefault) GetTicketAverageByDestinationCountry(country string) (float64, error) {
	percentage, err := s.rp.GetTicketAverageByDestinationCountry(country)

	return percentage, err
}
