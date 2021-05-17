package predictions

import (
	"encoding/json"
	"fmt"

	"github.com/webercoder/gocean/src/coops/coopsclient"
)

// NewAPI creates a new API for interacting retrieving tide predictions.
func NewAPI(app string) *API {
	return &API{
		Client: coopsclient.NewClient(app),
	}
}

// Retrieve gets the predictions from the station.
func (predAPI *API) Retrieve(
	station string,
	hours int,
) ([]Prediction, error) {
	jsonData, err := predAPI.Client.GetJSON(
		coopsclient.NewClientRequest(
			coopsclient.WithProduct(coopsclient.ProductPredictions),
			coopsclient.WithStation(station),
			coopsclient.WithHours(hours),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error reading predictions request body: %v", err)
	}

	predictions := &Predictions{}
	err = json.Unmarshal(jsonData, &predictions)
	if err != nil {
		return nil, fmt.Errorf("error parsing predictions json data: %v", err)
	}

	return predictions.Predictions, nil
}

// PrintTabDelimited outputs the tides in text format.
func (predAPI *API) PrintTabDelimited(station string, predictions []Prediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
