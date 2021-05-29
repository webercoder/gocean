package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// WaterTemperatureAPI interacts with the water temperature product.
type WaterTemperatureAPI struct {
	App    string
	Client *Client
}

// WaterTemperatureResult is a deserialized water temperature response.
type WaterTemperatureResult struct {
	XMLName           xml.Name           `xml:"data"`
	WaterTemperatures []WaterTemperature `xml:"observations" json:"data"`
}

// WaterTemperature is a singular, deserialized water temperature.
type WaterTemperature struct {
	XMLName xml.Name `xml:"at"`
	Time    string   `xml:"t,attr" json:"t"`
	Value   string   `xml:"v,attr" json:"v"`
}

// NewWaterTemperatureAPI creates a new water temperature API client.
func NewWaterTemperatureAPI(app string) *WaterTemperatureAPI {
	return &WaterTemperatureAPI{
		Client: NewClient(app),
	}
}

// GetWaterTemperatures gets the WaterTemperatures from the station.
func (api *WaterTemperatureAPI) GetWaterTemperatures(req *ClientRequest) ([]WaterTemperature, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading water temperature request body: %v", err)
	}

	levels := &WaterTemperatureResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing water temperature json data: %v", err)
	}

	if len(levels.WaterTemperatures) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing water temperature json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.WaterTemperatures, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *WaterTemperatureAPI) PrintTabDelimited(station string, temps []WaterTemperature) {
	fmt.Println("Water temperatures for station:", station)
	for _, t := range temps {
		fmt.Printf("  %s\t%s\n", t.Time, t.Value)
	}
}
