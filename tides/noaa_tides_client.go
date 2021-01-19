package tides

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/webercoder/gocean/utils"
)

// DefaultTidesEndpoint is the default tides endpoint from NOAA.
const DefaultTidesEndpoint = "https://api.tidesandcurrents.noaa.gov/api/prod/datagetter"

// https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?station=9410170&
// product=predictions&units=metric&time_zone=lst_ldt&application=gocean&format=json&
// datum=STND&begin_date=20210119&end_date=20210121

// NOAATidesClient interacts with the NOAA api.
type NOAATidesClient struct {
	Application string
	Client      utils.HTTPGetter
	URL         string
}

// NewNOAATidesClient creates a new NOAATidesClient object with default values.
func NewNOAATidesClient() *NOAATidesClient {
	return &NOAATidesClient{URL: DefaultTidesEndpoint, Client: &http.Client{}}
}

// RetrievePredictions gets the predictions from station.
func (ntc *NOAATidesClient) RetrievePredictions(station string, hours int) ([]TidePrediction, error) {
	currentTime := time.Now()

	baseURL, err := url.Parse(ntc.URL)
	if err != nil {
		fmt.Println("Unable to parse API URL", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("application", "gocean")
	params.Add("begin_date", currentTime.Format("20060102 15:04"))
	params.Add("end_date", currentTime.Add(time.Hour*time.Duration(hours)).Format("20060102 15:04"))
	params.Add("datum", "STND")
	params.Add("format", "json")
	params.Add("product", "predictions")
	params.Add("station", station)
	params.Add("time_zone", "lst_ldt")
	params.Add("units", "english")
	baseURL.RawQuery = params.Encode()

	resp, err := ntc.Client.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error retrieving predictions", err)
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading predictions request body", err)
		return nil, err
	}

	predictions := &TidePredictions{}
	err = json.Unmarshal(jsonData, &predictions)
	if err != nil {
		fmt.Println("Error parsing predictions json data", err)
		return nil, err
	}

	return predictions.Predictions, nil
}

// PrintPredictions outputs the tides in text format.
func PrintPredictions(station string, predictions []TidePrediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
