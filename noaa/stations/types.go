package stations

import (
	"encoding/xml"

	"github.com/webercoder/gocean/utils"
)

// StationDistance describes the distance from a set of coordinates to a station.
type StationDistance struct {
	Distance float64
	From     utils.GeoCoordinates
	Station  Station
}

// GetStationResponse represents a response to retrieving the station data
type GetStationResponse struct {
	XMLName  xml.Name  `xml:"stations"`
	Stations []Station `xml:"station"`
}

// Station represents a station
type Station struct {
	XMLName      xml.Name            `xml:"station"`
	Name         string              `xml:"name,attr"`
	ID           int                 `xml:"ID,attr"`
	Metadata     StationMetadata     `xml:"metadata"`
	Capabilities []StationCapability `xml:"parameter"`
}

// StationMetadata contains metadata about a station.
type StationMetadata struct {
	XMLName         xml.Name        `xml:"metadata"`
	Location        StationLocation `xml:"location"`
	DateEstablished string          `xml:"date_established"`
}

// StationLocation contains information about the station's location.
type StationLocation struct {
	XMLName xml.Name `xml:"location"`
	Lat     float64  `xml:"lat"`
	Long    float64  `xml:"long"`
	State   string   `xml:"state"`
}

// StationCapability specifies one NOAA capability for the NOAA API.
type StationCapability struct {
	XMLName  xml.Name `xml:"parameter"`
	Name     string   `xml:"name,attr"`
	SensorID string   `xml:"sensorID,attr"`
	DCP      int      `xml:"DCP,attr"`
	Status   int      `xml:"status,attr"`
}
