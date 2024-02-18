package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type BoundingBox struct {
	SouthWestLatitude  float64 `json:"south_west_latitude"`
	SouthWestLongitude float64 `json:"south_west_longitude"`
	NorthEastLatitude  float64 `json:"north_east_latitude"`
	NorthEastLongitude float64 `json:"north_east_longitude"`
}

type FlightFilter struct {
	BoundingBox  *BoundingBox `json:"bounding_box"`
	Zoom         int64        `json:"zoom"`
	Hex          string       `json:"hex"`
	AirlineICAO  string       `json:"airline_icao"`
	AirlineIATA  string       `json:"airline_iata"`
	Flag         string       `json:"flag"`
	FlightICAO   string       `json:"flight_icao"`
	FlightIATA   string       `json:"flight_iata"`
	FlightNumber string       `json:"flight_number"`
}

type Repository interface {
	GetOne(ctx context.Context, id string) (*Flight, error)
	GetMany(ctx context.Context, flightFilter *FlightFilter) ([]*Flight, error)
}

type repository struct {
	flights []*Flight
}

func (r *repository) GetOne(ctx context.Context, id string) (*Flight, error) {
	var filtersFn = []FilterFn{
		FilterByFlightICAO
		FilterByFlightIATA
	}

	for _, f := range r.flights {
		if FilterOr(f, flightFilter, filtersFn) {
			return f, nil
		}
	}

	return nil, fmt.Errorf("flight not found")
}

func (r *repository) GetMany(ctx context.Context, flightFilter *FlightFilter) ([]*Flight, error) {
	var flights []*Flight
	var filtersFn = []FilterFn{
		FilterByBoundingBoxOptional,
		FilterByZoomOptional,
		FilterByHexOptional,
		FilterByAirlineICAOOptional,
		FilterByAirlineIATAOptional,
		FilterByFlagOptional,
		FilterByFlightICAOOptional,
		FilterByFlightIATAOptional,
		FilterByFlightNumberOptional,
	}

	for _, f := range r.flights {
		if !FilterByAll(f, flightFilter, filtersFn) {
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

func FilterByAll(f *Flight, filter *FlightFilter, filtersFn []FilterFn) bool {
	for _, filterFn := range filtersFn {
		if !filterFn(f, filter) {
			return false
		}
	}
	return true
}

func FilterOr(f *Flight, filter *FlightFilter, filtersFn []FilterFn) bool {
	for _, filterFn := range filtersFn {
		if filterFn(f, filter) {
			return true
		}
	}
	return false
}
