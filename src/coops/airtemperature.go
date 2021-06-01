package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// AirTemperatureAPI interacts with the air temperature product.
type AirTemperatureAPI struct {
	App    string
	Client *Client
}

// AirTemperatureResult is a deserialized air temperature response.
type AirTemperatureResult struct {
	XMLName         xml.Name         `xml:"data"`
	AirTemperatures []AirTemperature `xml:"observations" json:"data"`
}

// AirTemperature is a singular, deserialized air temperature.
type AirTemperature struct {
	XMLName xml.Name `xml:"at"`
	ValueBasedResult
}

// NewAirTemperatureAPI creates a new air temperature API client.
func NewAirTemperatureAPI(app string) *AirTemperatureAPI {
	return &AirTemperatureAPI{
		Client: NewClient(app),
	}
}

// GetAirTemperatures gets the AirTemperatures from the station.
func (api *AirTemperatureAPI) GetAirTemperatures(req *ClientRequest) ([]AirTemperature, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading air temperature request body: %v", err)
	}

	levels := &AirTemperatureResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing air temperature json data: %v", err)
	}

	if len(levels.AirTemperatures) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing air temperature json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.AirTemperatures, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *AirTemperatureAPI) PrintTabDelimited(req *ClientRequest, temps []AirTemperature) {
	fmt.Printf("Air temperature readings for station %s:\n\n", req.Station)
	units := "f"
	if req.Units == UnitsMetric {
		units = "c"
	}
	for _, el := range temps {
		fmt.Printf("%s: %s°%s\n", el.Time, el.Value, units)
	}
}
