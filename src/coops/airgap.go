package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// AirGapAPI interacts with the wind product.
type AirGapAPI struct {
	App    string
	Client *Client
}

// AirGapResult is a deserialized air gap response.
type AirGapResult struct {
	XMLName xml.Name `xml:"data"`
	AirGaps []AirGap `xml:"observations" json:"data"`
}

// AirGap is a singular, deserialized air gap.
type AirGap struct {
	XMLName xml.Name `xml:"at"`
	ValueBasedResultWithFlags
	Sigma string `xml:"s,attr" json:"s"`
}

// NewAirGapAPI creates a new air gap API client.
func NewAirGapAPI(app string) *AirGapAPI {
	return &AirGapAPI{
		Client: NewClient(app),
	}
}

// GetAirGap gets the AirGaps from the station.
func (api *AirGapAPI) GetAirGap(req *ClientRequest) ([]AirGap, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading air gap request body: %v", err)
	}

	levels := &AirGapResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing air gap json data: %v", err)
	}

	if len(levels.AirGaps) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing air gap json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.AirGaps, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *AirGapAPI) PrintTabDelimited(req *ClientRequest, airgaps []AirGap) {
	fmt.Printf("Air gap readings for station %s:\n\n", req.Station)
	units := "ft"
	if req.Units == UnitsMetric {
		units = "m"
	}
	for _, el := range airgaps {
		fmt.Printf("%s: %s%s (Sigma: %s, Flags: %s)\n", el.Time, el.Value, units, el.Sigma, el.Flags)
	}
}
