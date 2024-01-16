package repository

import (
	"app/internal"
	"context"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketByDestinationCountry(country string) map[int]internal.TicketAttributes {
	// create a copy of the map
	t := make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return t
}

func (r *RepositoryTicketMap) GetTicketAverageByDestinationCountry(country string) (float64, error) {
	// create a copy of the map
	count := 0
	for _, v := range r.db {
		if v.Country == country {
			count++
		}
	}

	if count == 0 {
		err := internal.ErrCountryNotFound
		return 0, err
	}

	return float64(count) / float64(r.lastId) * 100.0, nil
}
