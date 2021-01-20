package waterlevels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/webercoder/gocean/noaa/tidesandcurrents"
)

// Predictions contain tide predictions from a station
type predictions struct {
	Predictions []Prediction `json:"predictions"`
}

// Prediction contains a single tide prediction for a specific time.
type Prediction struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

// RetrievePredictions gets the predictions from station.
//
// Example query:
// https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?station=9410170&
// product=predictions&units=metric&time_zone=lst_ldt&application=gocean&format=json&
// datum=STND&begin_date=20210119&end_date=20210121
func RetrievePredictions(
	ntc *tidesandcurrents.Client,
	station string,
	hours int,
) ([]Prediction, error) {
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
	params.Add("datum", "MLLW")
	params.Add("format", "json")
	params.Add("product", "predictions")
	params.Add("station", station)
	params.Add("time_zone", "lst_ldt")
	params.Add("units", "english")
	baseURL.RawQuery = params.Encode()

	resp, err := ntc.HTTPClient.Get(baseURL.String())
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

	predictions := &predictions{}
	err = json.Unmarshal(jsonData, &predictions)
	if err != nil {
		fmt.Println("Error parsing predictions json data", err)
		return nil, err
	}

	return predictions.Predictions, nil
}

// PrintPredictions outputs the tides in text format.
func PrintPredictions(station string, predictions []Prediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
