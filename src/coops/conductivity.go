package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ConductivityAPI interacts with the wind product.
type ConductivityAPI struct {
	App    string
	Client *Client
}

// ConductivityResult is a deserialized conductivity response.
type ConductivityResult struct {
	XMLName       xml.Name       `xml:"data"`
	Conductivitys []Conductivity `xml:"observations" json:"data"`
}

// Conductivity is a singular, deserialized conductivity.
type Conductivity struct {
	XMLName xml.Name `xml:"ct"`
	ValueBasedResultWithFlags
}

// NewConductivityAPI creates a new conductivity API client.
func NewConductivityAPI(app string) *ConductivityAPI {
	return &ConductivityAPI{
		Client: NewClient(app),
	}
}

// GetConductivity gets the Conductivitys from the station.
func (api *ConductivityAPI) GetConductivity(req *ClientRequest) ([]Conductivity, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading conductivity request body: %v", err)
	}

	levels := &ConductivityResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing conductivity json data: %v", err)
	}

	if len(levels.Conductivitys) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing conductivity json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.Conductivitys, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *ConductivityAPI) PrintTabDelimited(req *ClientRequest, conductivityData []Conductivity) {
	fmt.Printf("Conductivity readings for station %s:\n\n", req.Station)
	for _, el := range conductivityData {
		fmt.Printf("%s: %smS/cm (Flags: %s)\n", el.Time, el.Value, el.Flags)
	}
}
