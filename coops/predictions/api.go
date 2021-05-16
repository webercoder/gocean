package predictions

import (
	"encoding/json"
	"fmt"

	"github.com/webercoder/gocean/coops/coops_client"
)

func NewPredictionsAPI(app string) *PredictionsAPI {
	return &PredictionsAPI{
		Client: coops_client.NewClient(app),
	}
}

// Retrieve gets the predictions from the station.
func (predAPI *PredictionsAPI) Retrieve(
	station string,
	hours int,
) ([]Prediction, error) {
	jsonData, err := predAPI.Client.GetJSON(
		coops_client.NewClientRequest(
			coops_client.WithProduct(coops_client.ProductPredictions),
			coops_client.WithStation(station),
			coops_client.WithHours(hours),
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
func (predAPI *PredictionsAPI) PrintTabDelimited(station string, predictions []Prediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
