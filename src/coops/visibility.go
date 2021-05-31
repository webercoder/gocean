package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// VisibilityAPI interacts with the wind product.
type VisibilityAPI struct {
	App    string
	Client *Client
}

// VisibilityResult is a deserialized visibility response.
type VisibilityResult struct {
	XMLName     xml.Name     `xml:"data"`
	Visibilitys []Visibility `xml:"observations" json:"data"`
}

// Visibility is a singular, deserialized visibility.
type Visibility struct {
	XMLName xml.Name `xml:"vi"`
	ValueBasedResultWithFlags
}

// NewVisibilityAPI creates a new visibility API client.
func NewVisibilityAPI(app string) *VisibilityAPI {
	return &VisibilityAPI{
		Client: NewClient(app),
	}
}

// GetVisibility gets the Visibilitys from the station.
func (api *VisibilityAPI) GetVisibility(req *ClientRequest) ([]Visibility, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading visibility request body: %v", err)
	}

	levels := &VisibilityResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing visibility json data: %v", err)
	}

	if len(levels.Visibilitys) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing visibility json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.Visibilitys, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *VisibilityAPI) PrintTabDelimited(station string, visibilityData []Visibility) {
	fmt.Println("Visibility readings for station:", station)
	for _, el := range visibilityData {
		fmt.Printf("\t%s\t%s (%s)\n", el.Time, el.Value, el.Flags)
	}
}
