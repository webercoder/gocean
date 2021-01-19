package stations

import "io"

// StationRetriever .
type StationRetriever interface {
	GetStations(url string) *io.Writer
}
