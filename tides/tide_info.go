package tides

// TideInfo Contains tide information for a single location
// Datatypes match up with examples on https://api.tidesandcurrents.noaa.gov/api/prod/
type TideInfo struct {
	metadata TideInfoMetadata
	data     []TideInfoData
}

// TideInfoMetadata contains information about the measuring station.
type TideInfoMetadata struct {
	id   int
	name string
	lat  float64
	lon  float64
}

// TideInfoData contains the actual tide data from the station
type TideInfoData struct {
	t string
	v string
	f string
	q string
}
