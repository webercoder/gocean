package water_level

import (
	"encoding/json"
	"fmt"

	"github.com/webercoder/gocean/coops/coops_client"
)

func NewWaterLevelAPI(app string) *WaterLevelAPI {
	return &WaterLevelAPI{
		Client: coops_client.NewClient(app),
	}
}

// Retrieve gets the WaterLevels from the station.
func (predAPI *WaterLevelAPI) Retrieve(
	station string,
	hours int,
) ([]WaterLevel, error) {
	jsonData, err := predAPI.Client.GetJSON(
		coops_client.NewClientRequest(
			coops_client.WithProduct(coops_client.ProductWaterLevel),
			coops_client.WithStation(station),
			coops_client.WithHours(hours),
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
		jsonErrResp := &coops_client.ClientErrorResponse{}
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
func (predAPI *WaterLevelAPI) PrintTabDelimited(station string, WaterLevels []WaterLevel) {
	fmt.Println("Tide water levels for station:", station)
	for _, WaterLevel := range WaterLevels {
		fmt.Printf("  %s\t%s\t%s\n", WaterLevel.Time, WaterLevel.Value, WaterLevel.Quality)
	}
}
