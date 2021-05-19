package coops

import (
	"encoding/json"
	"fmt"
)

// PredictionsResult contains tide predictions from a station
type PredictionsResult struct {
	Predictions []Prediction `json:"predictions"`
}

// Prediction contains a single tide prediction for a specific time.
type Prediction struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

// PredictionsAPI interacts with the predictions product.
type PredictionsAPI struct {
	App    string
	Client *Client
}

// NewPredictionsAPI creates a new API for interacting retrieving tide predictions.
func NewPredictionsAPI(app string) *PredictionsAPI {
	return &PredictionsAPI{
		Client: NewClient(app),
	}
}

// Retrieve gets the predictions from the station.
func (api *PredictionsAPI) Retrieve(
	station string,
	hours int,
) ([]Prediction, error) {
	jsonData, err := api.Client.GetJSON(
		NewClientRequest(
			WithProduct(ProductPredictions),
			WithStation(station),
			WithHours(hours),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error reading predictions request body: %v", err)
	}

	predictions := &PredictionsResult{}
	err = json.Unmarshal(jsonData, &predictions)
	if err != nil {
		return nil, fmt.Errorf("error parsing predictions json data: %v", err)
	}

	return predictions.Predictions, nil
}

// PrintTabDelimited outputs the tides in text format.
func (api *PredictionsAPI) PrintTabDelimited(station string, predictions []Prediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
