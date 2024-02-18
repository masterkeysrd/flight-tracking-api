package flight

import "context"

type GetFlightDataRequest struct {
	BoundingBox        []float64 `json:"bbox"`
	Zoom               int64     `json:"zoom"`
	Hex                string    `json:"hex"`
	RegistrationNumber string    `json:"reg_number"`
	AirlineICAO        string    `json:"airline_icao"`
	AirlineIATA        string    `json:"airline_iata"`
	Flag               string    `json:"flag"`
	FlightICAO         string    `json:"flight_icao"`
	FlightIATA         string    `json:"flight_iata"`
	FlightNumber       string    `json:"flight_number"`
}

type GetFlightDataResponse struct {
	Data []*GetFlightDataResponseItem `json:"data"`
}

type GetFlightDataResponseItem struct {
	Hex                string  `json:"hex"`
	RegistrationNumber string  `json:"reg_number"`
	Flag               string  `json:"flag"`
	Latitude           float64 `json:"lat"`
	Longitude          float64 `json:"lgn"`
	Altitude           int64   `json:"alt"`
	Direction          int64   `json:"dir"`
	Speed              int64   `json:"speed"`
	Velocity           int64   `json:"v_speed"`
	Squawk             string  `json:"squawk"`
	FlightNumber       string  `json:"flight_number"`
	FlightICAO         string  `json:"flight_icao"`
	FlightIATA         string  `json:"flight_iata"`
	AirlineICAO        string  `json:"airline_icao"`
	AirlineIATA        string  `json:"airline_iata"`
	AircraftICAO       string  `json:"aircraft_icao"`
	Updated            int64   `json:"updated"`
	Status             string  `json:"status"`
}

type SearchFlightInfoRequest struct {
	FlightICAO string `json:"flight_icao"`
	FlightIATA string `json:"flight_iata"`
}

type SearchFlightInfoResponse struct {
	Hex                string  `json:"hex"`
	RegistrationNumber string  `json:"reg_number"`
	AircraftICAO       string  `json:"aircraft_icao"`
	Flag               string  `json:"flag"`
	Latitude           float64 `json:"lat"`
	Longitude          float64 `json:"lgn"`
	Altitude           int64   `json:"alt"`
	Direction          int64   `json:"dir"`
	Speed              int64   `json:"speed"`
	Velocity           int64   `json:"v_speed"`
	Squawk             string  `json:"squawk"`
	AirlineICAO        string  `json:"airline_icao"`
	AirlineIATA        string  `json:"airline_iata"`
	FlightNumber       string  `json:"flight_number"`
	FlightICAO         string  `json:"flight_icao"`
	FlightIATA         string  `json:"flight_iata"`
	Duration           int64   `json:"duration"`
	Updated            int64   `json:"updated"`
	Status             string  `json:"status"`
}

type Service interface {
	GetFlightData(ctx context.Context, req *GetFlightDataRequest) (*GetFlightDataResponse, error)
	SearchFlightInfo(ctx context.Context, req *SearchFlightInfoRequest) (*SearchFlightInfoResponse, error)
}

type service struct {
	repository Repository
	mapper     Mapper
}

func NewService(repository Repository, mapper Mapper) Service {
	return &service{
		repository: repository,
		mapper:     mapper,
	}
}

func (s *service) GetFlightData(ctx context.Context, req *GetFlightDataRequest) (*GetFlightDataResponse, error) {
	filter := s.mapper.FromGetFlightDataRequestToFlightFilterParams(req)

	flights, err := s.repository.GetMany(ctx, filter)

	if err != nil {
		return nil, err
	}

	return s.mapper.FromFlightToGetFlightDataResponse(flights), nil
}

func (s *service) SearchFlightInfo(ctx context.Context, req *SearchFlightInfoRequest) (*SearchFlightInfoResponse, error) {
	filter := s.mapper.FromSearchFlightInfoRequestToFlightFilterParams(req)

	f, err := s.repository.GetOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return s.mapper.FromFlightToSearchFlightInfoResponse(f), nil
}
