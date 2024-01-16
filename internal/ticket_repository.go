package internal

import "errors"

var (
	ErrCountryNotFound = errors.New("country not found")
)

// TicketRepository represents the repository interface for tickets
type TicketRepository interface {
	// Get returns all the tickets
	// Get(ctx context.Context) (t map[int]TicketAttributes, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketByDestinationCountry(country string) map[int]TicketAttributes
}
