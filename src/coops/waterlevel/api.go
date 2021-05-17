package waterlevel

import (
	"encoding/json"
	"fmt"

	"github.com/webercoder/gocean/src/coops/coopsclient"
)

// NewAPI creates a new water level API client.
func NewAPI(app string) *API {
	return &API{
		Client: coopsclient.NewClient(app),
	}
}

// Retrieve gets the WaterLevels from the station.
func (predAPI *API) Retrieve(
	station string,
	hours int,
) ([]WaterLevel, error) {
	jsonData, err := predAPI.Client.GetJSON(
		coopsclient.NewClientRequest(
			coopsclient.WithProduct(coopsclient.ProductWaterLevel),
			coopsclient.WithStation(station),
			coopsclient.WithHours(hours),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error reading waterlevels request body: %v", err)
	}

	levels := &WaterLevels{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing waterlevels json data: %v", err)
	}

	if len(levels.WaterLevels) == 0 {
		jsonErrResp := &coopsclient.ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing waterlevels json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.WaterLevels, nil
}

// PrintTabDelimited outputs the tides in text format.
func (predAPI *API) PrintTabDelimited(station string, levels []WaterLevel) {
	fmt.Println("Tide water levels for station:", station)
	for _, level := range levels {
		quality := "Preliminary"
		if level.Quality == "v" {
			quality = "Verified"
		}

		fmt.Printf("  %s\t%s\t%s\n", level.Time, level.Value, quality)
	}
}
