package stations

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/webercoder/gocean/utils"
	"golang.org/x/net/html/charset"
)

// DefaultStationsEndpoint is the default stations endpoint from NOAA.
const DefaultStationsEndpoint = "https://opendap.co-ops.nos.noaa.gov/stations/stationsXML.jsp"

// Client interacts with NOAA.
type Client struct {
	URL           string
	HTTPClient    utils.HTTPGetter
	StationsCache []Station
}

// NewClient creates a new NOAAStationClient with the default URL.
func NewClient() *Client {
	return &Client{URL: DefaultStationsEndpoint, HTTPClient: &http.Client{}}
}

func decodeStationsXMLStream(httpResponseBody io.ReadCloser) (GetStationResponse, error) {
	var stations GetStationResponse
	decoder := xml.NewDecoder(httpResponseBody)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&stations)
	return stations, err
}

// GetStations .
func (s *Client) GetStations(skipCache bool) []Station {
	var stationResponse GetStationResponse

	if !skipCache && len(s.StationsCache) > 0 {
		return s.StationsCache
	}

	resp, err := s.HTTPClient.Get(s.URL)
	if err != nil {
		fmt.Println("Error retrieving station data", err)
		return stationResponse.Stations
	}
	if resp.Body == nil {
		fmt.Println("Error retrieving HTTP request body for station data", err)
		return stationResponse.Stations
	}
	defer resp.Body.Close()

	stationResponse, err = decodeStationsXMLStream(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	s.StationsCache = stationResponse.Stations
	return s.StationsCache
}

// GetNearestStation gets the nearest station to a given set of coordinates.
func (s *Client) GetNearestStation(coords utils.GeoCoordinates) (Station, float64) {
	var nearestStation Station
	var nearestDistance float64 = -1.0

	for _, station := range s.GetStations(false) {
		distance := utils.HarvesineDistance(coords, utils.GeoCoordinates{
			Lat:  station.Metadata.Location.Lat,
			Long: station.Metadata.Location.Long,
		})
		if nearestDistance < 0 || distance < nearestDistance {
			nearestStation = station
			nearestDistance = distance
		}
	}

	return nearestStation, nearestDistance
}