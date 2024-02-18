package flight

type Flight struct {
	Hex            string `json:"hex"`
	RegisterNumber string `json:"reg_number"`
	AircraftICAO   string `json:"aircraft_icao"`
	Flag           string `json:"flag"`
	Latitude       float  `json:"lat"`
	Longitude      float  `json:"lgn"`
	Altitude       int64  `json:"alt"`
	Direction      int64  `json:"dir"`
	Speed          int64  `json:"speed"`
	Squawk         string `json:"squawk"`
	AirlineICAO    string `json:"airline_icao"`
	AirlineIATA    string `json:"airline_iata"`
	FlightNumber   string `json:"flight_number"`
	FlightICAO     string `json:"flight_icao"`
	FlightIATA     string `json:"flight_iata"`
	Duration       string `json:"duration"`
	Updated        int64  `json:"updated"`
	Status         string `json:"status"`
}

func (f *Flight) IsInBoundingBox(southWestLatitude, southWestLongitude, northEastLatitude, northEastLongitude float64) bool {
	return f.Latitude >= southWestLatitude && f.Latitude <= northEastLatitude &&
		f.Longitude >= southWestLongitude && f.Longitude <= northEastLongitude
}
