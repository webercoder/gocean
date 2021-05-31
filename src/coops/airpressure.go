package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// AirPressureAPI interacts with the wind product.
type AirPressureAPI struct {
	App    string
	Client *Client
}

// AirPressureResult is a deserialized air pressure response.
type AirPressureResult struct {
	XMLName      xml.Name      `xml:"data"`
	AirPressures []AirPressure `xml:"observations" json:"data"`
}

// AirPressure is a singular, deserialized air pressure.
type AirPressure struct {
	XMLName xml.Name `xml:"at"`
	Time    string   `xml:"t,attr" json:"t"`
	Value   string   `xml:"v,attr" json:"v"`
	F       string   `xml:"f,attr" json:"f"`
}

// NewAirPressureAPI creates a new air pressure API client.
func NewAirPressureAPI(app string) *AirPressureAPI {
	return &AirPressureAPI{
		Client: NewClient(app),
	}
}

// GetAirPressure gets the AirPressures from the station.
func (api *AirPressureAPI) GetAirPressure(req *ClientRequest) ([]AirPressure, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading air pressure request body: %v", err)
	}

	levels := &AirPressureResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing air pressure json data: %v", err)
	}

	if len(levels.AirPressures) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing air pressure json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.AirPressures, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *AirPressureAPI) PrintTabDelimited(station string, airpressures []AirPressure) {
	fmt.Println("Air pressure readings for station:", station)
	for _, ap := range airpressures {
		fmt.Printf("\t%s\t%s\n", ap.Time, ap.Value)
	}
}
