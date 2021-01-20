package stations

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"

	"github.com/webercoder/gocean/utils"
	"golang.org/x/net/html/charset"
)

// DefaultStationsEndpoint is the default stations endpoint from NOAA.
const DefaultStationsEndpoint = "https://opendap.co-ops.nos.noaa.gov/stations/stationsXML.jsp"

// MaxStationsSearchChunkSize is used when searching for nearest station with goroutines
const MaxStationsSearchChunkSize = 10

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
func (s *Client) GetNearestStation(coords utils.GeoCoordinates) *StationDistance {
	result := &StationDistance{Distance: -1.0, From: coords}
	stations := s.GetStations(false)
	var wg sync.WaitGroup

	if len(stations) == 0 {
		return result
	}

	// Either MaxStationsSearchChunkSize or half the size of the stations list rounded up
	chunkSize := math.Min(MaxStationsSearchChunkSize, math.Ceil(float64(len(stations))/2))

	// Length is the number of goroutines that will be spawned
	routineCount := int(math.Ceil(float64(len(stations)) / chunkSize))
	c := make(chan *StationDistance, routineCount)

	for i := 0; i < routineCount; i++ {
		start := int(float64(i) * chunkSize)
		end := int(math.Min(float64(len(stations)), float64(i+1)*chunkSize))
		wg.Add(1)
		go s.findNearestStation(&wg, c, coords, stations[start:end])
	}

	wg.Wait()
	close(c)
	for item := range c {
		if result.Distance == -1 || result.Distance > item.Distance {
			result = item
		}
	}

	return result
}

func (s *Client) findNearestStation(
	wg *sync.WaitGroup,
	c chan *StationDistance,
	coords utils.GeoCoordinates,
	stations []Station,
) {
	defer wg.Done()
	result := &StationDistance{Distance: -1.0, From: coords}

	for _, station := range stations {
		distance := utils.HarvesineDistance(coords, utils.GeoCoordinates{
			Lat:  station.Metadata.Location.Lat,
			Long: station.Metadata.Location.Long,
		})

		if result.Distance < 0 || result.Distance > distance {
			result.Station = station
			result.Distance = distance
		}
	}

	c <- result
}
