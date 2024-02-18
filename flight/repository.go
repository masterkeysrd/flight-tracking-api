package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type BoundingBox struct {
	SouthWestLatitude  float64
	SouthWestLongitude float64
	NorthEastLatitude  float64
	NorthEastLongitude float64
}

type FlightFilterParams struct {
	BoundingBox        *BoundingBox
	Zoom               int64
	Hex                string
	RegistrationNumber string
	AirlineICAO        string
	AirlineIATA        string
	Flag               string
	FlightICAO         string
	FlightIATA         string
	FlightNumber       string
}

type Repository interface {
	GetOne(ctx context.Context, params *FlightFilterParams) (*Flight, error)
	GetMany(ctx context.Context, params *FlightFilterParams) ([]*Flight, error)
}

type repository struct {
	flights []*Flight
}

func (r *repository) GetOne(ctx context.Context, params *FlightFilterParams) (*Flight, error) {
	log.Printf("flightRepository [GetOne]: %+v", params)
	filtersFn := []FilterFn{
		FilterByFlightICAO,
		FilterByFlightIATA,
	}

	if params == nil || (params.FlightICAO == "" && params.FlightIATA == "") {
		return nil, fmt.Errorf("flight_icao or flight_iata is required")
	}

	for _, f := range r.flights {
		if FilterOr(f, params, filtersFn) {
			return f, nil
		}
	}

	return nil, fmt.Errorf("flight not found")
}

func (r *repository) GetMany(ctx context.Context, params *FlightFilterParams) ([]*Flight, error) {
	log.Printf("flightRepository [GetMany]: %+v", params)
	var flights []*Flight
	var filtersFn = []FilterFn{
		FilterByBoundingBoxOptional,
		FilterByZoomOptional,
		FilterByHexOptional,
		FilterByRegistrationNumberOptional,
		FilterByAirlineICAOOptional,
		FilterByAirlineIATAOptional,
		FilterByFlagOptional,
		FilterByFlightICAOOptional,
		FilterByFlightIATAOptional,
		FilterByFlightNumberOptional,
	}

	for _, f := range r.flights {
		if !FilterByAll(f, params, filtersFn) {
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

func FilterByAll(f *Flight, params *FlightFilterParams, filtersFn []FilterFn) bool {
	for _, filterFn := range filtersFn {
		if !filterFn(f, params) {
			return false
		}
	}
	return true
}

func FilterOr(f *Flight, params *FlightFilterParams, filtersFn []FilterFn) bool {
	for _, filterFn := range filtersFn {
		if filterFn(f, params) {
			return true
		}
	}
	return false
}
