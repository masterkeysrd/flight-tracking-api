package flight

type Mapper interface {
	FromGetFlightDataRequestToFlightFilterParams(req *GetFlightDataRequest) *FlightFilterParams
	FromSearchFlightInfoRequestToFlightFilterParams(req *SearchFlightInfoRequest) *FlightFilterParams
	FromFlightToGetFlightDataResponse(flights []*Flight) *GetFlightDataResponse
	FromFlightToSearchFlightInfoResponse(f *Flight) *SearchFlightInfoResponse
}

type mapper struct{}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) FromGetFlightDataRequestToFlightFilterParams(req *GetFlightDataRequest) *FlightFilterParams {
	var boundingBox *BoundingBox

	if len(req.BoundingBox) == 4 {
		boundingBox = &BoundingBox{
			SouthWestLatitude:  req.BoundingBox[0],
			SouthWestLongitude: req.BoundingBox[1],
			NorthEastLatitude:  req.BoundingBox[2],
			NorthEastLongitude: req.BoundingBox[3],
		}
	}

	return &FlightFilterParams{
		BoundingBox:        boundingBox,
		Zoom:               req.Zoom,
		Hex:                req.Hex,
		RegistrationNumber: req.RegistrationNumber,
		AirlineICAO:        req.AirlineICAO,
		AirlineIATA:        req.AirlineIATA,
		Flag:               req.Flag,
		FlightICAO:         req.FlightICAO,
		FlightIATA:         req.FlightIATA,
		FlightNumber:       req.FlightNumber,
	}
}

func (m *mapper) FromSearchFlightInfoRequestToFlightFilterParams(req *SearchFlightInfoRequest) *FlightFilterParams {
	return &FlightFilterParams{
		FlightICAO: req.FlightICAO,
		FlightIATA: req.FlightIATA,
	}
}

func (m *mapper) FromFlightToGetFlightDataResponse(flights []*Flight) *GetFlightDataResponse {
	var data []*GetFlightDataResponseItem
	for _, f := range flights {
		data = append(data, &GetFlightDataResponseItem{
			Hex:                f.Hex,
			RegistrationNumber: f.RegistrationNumber,
			AircraftICAO:       f.AircraftICAO,
			Flag:               f.Flag,
			Latitude:           f.Latitude,
			Longitude:          f.Longitude,
			Altitude:           f.Altitude,
			Direction:          f.Direction,
			Speed:              f.Speed,
			Velocity:           f.Velocity,
			Squawk:             f.Squawk,
			AirlineICAO:        f.AirlineICAO,
			AirlineIATA:        f.AirlineIATA,
			FlightICAO:         f.FlightICAO,
			FlightIATA:         f.FlightIATA,
			Updated:            f.Updated,
			Status:             f.Status,
		})
	}

	return &GetFlightDataResponse{Data: data}
}

func (m *mapper) FromFlightToSearchFlightInfoResponse(f *Flight) *SearchFlightInfoResponse {
	return &SearchFlightInfoResponse{
		Hex:                f.Hex,
		RegistrationNumber: f.RegistrationNumber,
		AircraftICAO:       f.AircraftICAO,
		Flag:               f.Flag,
		Latitude:           f.Latitude,
		Longitude:          f.Longitude,
		Altitude:           f.Altitude,
		Direction:          f.Direction,
		Speed:              f.Speed,
		Velocity:           f.Velocity,
		Squawk:             f.Squawk,
		AirlineICAO:        f.AirlineICAO,
		AirlineIATA:        f.AirlineIATA,
		FlightICAO:         f.FlightICAO,
		FlightIATA:         f.FlightIATA,
		Duration:           f.Duration,
		Updated:            f.Updated,
		Status:             f.Status,
	}
}
