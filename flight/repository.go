package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type BoundingBox struct {
	SouthWestLatitude  float64
	SouthWestLongitude float64
	NorthEastLatitude  float64
	NorthEastLongitude float64
}

type FlightFilter struct {
	BoundingBox    *BoundingBox
	Zoom           int64
	Hex            string
	RegisterNumber string
	AirlineICAO    string
	AirlineIATA    string
	Flag           string
	FlightICAO     string
	FlightIATA     string
}

type Repository interface {
	// GetMany returns a list of parameters
	GetMany(ctx context.Context, flightFilter *FlightFilter) ([]*Flight, error)
}

type repository struct {
	flights []*Flight
}

func (r *repository) GetMany(ctx context.Context, flightFilter *FlightFilter) ([]*Flight, error) {
	var flights []*Flight
	for _, f := range r.flights {
		if flightFilter.BoundingBox != nil && !ByBoundingBox(flightFilter, f) {
			continue
		}
		flights = append(flights, f)
	}
	return flights, nil
}

func NewRepository(config *Config) (Repository, error) {
	filename := config.Filename()

	if filename == "" {
		return nil, fmt.Errorf("flight data file not found")
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error opening flight data file: %v", err)
	}

	var flights []*Flight
	if err := json.NewDecoder(file).Decode(&flights); err != nil {
		return nil, fmt.Errorf("error decoding flight data: %v", err)
	}

	return &repository{
		flights: flights,
	}, nil
}
