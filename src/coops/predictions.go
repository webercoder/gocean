package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// PredictionsResult contains tide predictions from a station
type PredictionsResult struct {
	XMLName     xml.Name     `xml:"data"`
	Predictions []Prediction `json:"predictions"`
}

// Prediction contains a single tide prediction for a specific time.
type Prediction struct {
	XMLName xml.Name `xml:"pr"`
	ValueBasedResult
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

// GetPredictions gets the predictions from the station.
func (api *PredictionsAPI) GetPredictions(req *ClientRequest) ([]Prediction, error) {
	data, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading predictions request body: %v", err)
	}
	predictions := &PredictionsResult{}
	err = json.Unmarshal(data, &predictions)
	if err != nil {
		return nil, fmt.Errorf("error parsing predictions data: %v", err)
	}

	if len(predictions.Predictions) == 0 {
		errResp := &ClientErrorResponse{}
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing water level data: %v", err)
		}

		if errResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", errResp.Err.Message)
		}
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
