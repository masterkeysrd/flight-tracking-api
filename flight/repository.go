package flight

import (
	"context"
	"fmt"
	"time"
)

type BoundingBox struct {
	SouthWestLatitude  float64
	SouthWestLongitude float64
	NorthEastLatitude  float64
	NorthEastLongitude float64
}

type FlightFilter struct {
	BoundingBox *BoundingBox
	Zoom		int64
	Hex         string
	RegisterNumber string
	AirlineICAO string
	AirlineIATA string
	Flag         string
	FlightICAO    string
	FlightIATA   string
}

type interface Repository {
	// GetMany returns a list of parameters
	GetMany(ctx context.Context, flightFilter *FlightFilter) ([]*Flight, error)
}

type repository struct {
	flights []*Flight
}

func NewRepository() Repository {
	return &repository{
		flights: make([]*Flight, 0),
	}
}
