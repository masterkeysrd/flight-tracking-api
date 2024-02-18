package flight

type FilterFn func(f *Flight, filter *FlightFilter) bool

func FilterByBoundingBoxOptional(f *Flight, filter *FlightFilter) bool {
	if filter.BoundingBox == nil {
		return true
	}

	if filter.BoundingBox.SouthWestLatitude == 0 && filter.BoundingBox.SouthWestLongitude == 0 &&
		filter.BoundingBox.NorthEastLatitude == 0 && filter.BoundingBox.NorthEastLongitude == 0 {
		return true
	}

	return f.IsInBoundingBox(
		filter.BoundingBox.SouthWestLatitude,
		filter.BoundingBox.SouthWestLongitude,
		filter.BoundingBox.NorthEastLatitude,
		filter.BoundingBox.NorthEastLongitude,
	)
}

func FilterByZoomOptional(f *Flight, filter *FlightFilter) bool {
	return true
}

func FilterByHexOptional(f *Flight, filter *FlightFilter) bool {
	if filter.Hex == "" {
		return true
	}

	return f.Hex == filter.Hex
}

func FilterByAirlineICAOOptional(f *Flight, filter *FlightFilter) bool {
	if filter.AirlineICAO == "" {
		return true
	}

	return f.AirlineICAO == filter.AirlineICAO
}

func FilterByAirlineIATAOptional(f *Flight, filter *FlightFilter) bool {
	if filter.AirlineIATA == "" {
		return true
	}

	return f.AirlineIATA == filter.AirlineIATA
}

func FilterByFlagOptional(f *Flight, filter *FlightFilter) bool {
	if filter.Flag == "" {
		return true
	}

	return f.Flag == filter.Flag
}

func FilterByFlightICAOOptional(f *Flight, filter *FlightFilter) bool {
	if filter.FlightICAO == "" {
		return true
	}

	return f.FlightICAO == filter.FlightICAO
}

func FilterByFlightIATAOptional(f *Flight, filter *FlightFilter) bool {
	if filter.FlightIATA == "" {
		return true
	}

	return f.FlightIATA == filter.FlightIATA
}

func FilterByFlightNumberOptional(f *Flight, filter *FlightFilter) bool {
	if filter.FlightNumber == "" {
		return true
	}

	return f.FlightNumber == filter.FlightNumber
}

func FilterByFlightICAO(f *Flight, filter *FlightFilter) bool {
	return f.FlightICAO == filter.FlightICAO
}

func FilterByFlightIATA(f *Flight, filter *FlightFilter) bool {
	return f.FlightIATA == filter.FlightIATA
}
