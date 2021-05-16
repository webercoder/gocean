package predictions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/webercoder/gocean/noaa/tidesandcurrents/lib"
)

func NewPredictionApi(app string) *PredictionsApi {
	return &PredictionsApi{
		Client: lib.NewClient(app),
	}
}

// Retrieve gets the predictions from the station.
func (predApi *PredictionsApi) Retrieve(
	station string,
	hours int,
) ([]Prediction, error) {
	req := lib.NewClientRequest(lib.WithStation(station))
	resp, err := predApi.Client.Get(req)
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

// PrintTabDelimited outputs the tides in text format.
func (predApi *PredictionsApi) PrintTabDelimited(station string, predictions []Prediction) {
	fmt.Println("Tide predictions for station:", station)
	for _, prediction := range predictions {
		fmt.Printf("  %s\t%s\n", prediction.Time, prediction.Value)
	}
}
