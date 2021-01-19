package tides

// TidePredictions contain tide predictions from a station
type TidePredictions struct {
	Predictions []TidePrediction `json:"predictions"`
}

// TidePrediction contains a single tide prediction for a specific time.
type TidePrediction struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

// TideInfo Contains tide information from a station
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
