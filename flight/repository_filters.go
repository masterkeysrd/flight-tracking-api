package flight

type FilterFn func(f *Flight, params *FlightFilterParams) bool

func FilterByBoundingBoxOptional(f *Flight, params *FlightFilterParams) bool {
	if params.BoundingBox == nil {
		return true
	}

	if params.BoundingBox.SouthWestLatitude == 0 && params.BoundingBox.SouthWestLongitude == 0 &&
		params.BoundingBox.NorthEastLatitude == 0 && params.BoundingBox.NorthEastLongitude == 0 {
		return true
	}

	return f.IsInBoundingBox(
		params.BoundingBox.SouthWestLatitude,
		params.BoundingBox.SouthWestLongitude,
		params.BoundingBox.NorthEastLatitude,
		params.BoundingBox.NorthEastLongitude,
	)
}

func FilterByZoomOptional(f *Flight, params *FlightFilterParams) bool {
	return true
}

func FilterByHexOptional(f *Flight, params *FlightFilterParams) bool {
	if params.Hex == "" {
		return true
	}

	return f.Hex == params.Hex
}

func FilterByRegistrationNumberOptional(f *Flight, params *FlightFilterParams) bool {
	if params.RegistrationNumber == "" {
		return true
	}

	return f.RegistrationNumber == params.RegistrationNumber
}

func FilterByAirlineICAOOptional(f *Flight, params *FlightFilterParams) bool {
	if params.AirlineICAO == "" {
		return true
	}

	return f.AirlineICAO == params.AirlineICAO
}

func FilterByAirlineIATAOptional(f *Flight, params *FlightFilterParams) bool {
	if params.AirlineIATA == "" {
		return true
	}

	return f.AirlineIATA == params.AirlineIATA
}

func FilterByFlagOptional(f *Flight, params *FlightFilterParams) bool {
	if params.Flag == "" {
		return true
	}

	return f.Flag == params.Flag
}

func FilterByFlightICAOOptional(f *Flight, params *FlightFilterParams) bool {
	if params.FlightICAO == "" {
		return true
	}

	return f.FlightICAO == params.FlightICAO
}

func FilterByFlightIATAOptional(f *Flight, params *FlightFilterParams) bool {
	if params.FlightIATA == "" {
		return true
	}

	return f.FlightIATA == params.FlightIATA
}

func FilterByFlightNumberOptional(f *Flight, params *FlightFilterParams) bool {
	if params.FlightNumber == "" {
		return true
	}

	return f.FlightNumber == params.FlightNumber
}

func FilterByFlightICAO(f *Flight, params *FlightFilterParams) bool {
	return f.FlightICAO == params.FlightICAO
}

func FilterByFlightIATA(f *Flight, params *FlightFilterParams) bool {
	return f.FlightIATA == params.FlightIATA
}
