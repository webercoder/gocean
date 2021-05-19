package coops

import (
	"encoding/json"
	"fmt"
)

// WaterLevelAPI interacts with the water level product.
type WaterLevelAPI struct {
	App    string
	Client *Client
}

// WaterLevelResult is a deserialized water level response.
type WaterLevelResult struct {
	WaterLevels []WaterLevel `json:"data"`
}

// WaterLevel is a singular, deserialized water level.
type WaterLevel struct {
	Time    string `json:"t"`
	Value   string `json:"v"`
	Quality string `json:"q"`
}

// NewWaterLevelAPI creates a new water level API client.
func NewWaterLevelAPI(app string) *WaterLevelAPI {
	return &WaterLevelAPI{
		Client: NewClient(app),
	}
}

// Retrieve gets the WaterLevels from the station.
func (api *WaterLevelAPI) Retrieve(
	station string,
	hours int,
) ([]WaterLevel, error) {
	jsonData, err := api.Client.GetJSON(
		NewClientRequest(
			WithProduct(ProductWaterLevel),
			WithStation(station),
			WithHours(hours),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error reading water level request body: %v", err)
	}

	levels := &WaterLevelResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing water level json data: %v", err)
	}

	if len(levels.WaterLevels) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing water level json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.WaterLevels, nil
}

// PrintTabDelimited outputs the tides in text format.
func (api *WaterLevelAPI) PrintTabDelimited(station string, levels []WaterLevel) {
	fmt.Println("Tide water levels for station:", station)
	for _, level := range levels {
		quality := "Preliminary"
		if level.Quality == "v" {
			quality = "Verified"
		}

		fmt.Printf("  %s\t%s\t%s\n", level.Time, level.Value, quality)
	}
}
