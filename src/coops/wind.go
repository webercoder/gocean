package coops

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// WindAPI interacts with the wind product.
type WindAPI struct {
	App    string
	Client *Client
}

// WindResult is a deserialized wind response.
type WindResult struct {
	XMLName xml.Name `xml:"data"`
	Winds   []Wind   `xml:"observations" json:"data"`
}

// Wind is a singular, deserialized wind.
type Wind struct {
	XMLName          xml.Name `xml:"ws"`
	Time             string   `xml:"t,attr" json:"t"`
	Speed            string   `xml:"s,attr" json:"s"`
	DirectionDegrees string   `xml:"d,attr" json:"d"`
	DirectionAcronym string   `xml:"dr,attr" json:"dr"`
	Gusts            string   `xml:"g,attr" json:"g"`
	Flags            string   `xml:"f,attr" json:"f"`
}

// NewWindAPI creates a new wind API client.
func NewWindAPI(app string) *WindAPI {
	return &WindAPI{
		Client: NewClient(app),
	}
}

// GetWind gets the Winds from the station.
func (api *WindAPI) GetWind(req *ClientRequest) ([]Wind, error) {
	if req.Format != ResponseFormatJSON {
		req.Format = ResponseFormatJSON
	}

	jsonData, err := api.Client.Get(req)
	if err != nil {
		return nil, fmt.Errorf("error reading wind request body: %v", err)
	}

	levels := &WindResult{}
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, fmt.Errorf("error parsing wind json data: %v", err)
	}

	if len(levels.Winds) == 0 {
		jsonErrResp := &ClientErrorResponse{}
		err = json.Unmarshal(jsonData, &jsonErrResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing wind json data: %v", err)
		}

		if jsonErrResp.Err.Message != "" {
			return nil, fmt.Errorf("received error from API: %s", jsonErrResp.Err.Message)
		}
	}

	return levels.Winds, nil
}

// PrintTabDelimited outputs the data in text format.
func (api *WindAPI) PrintTabDelimited(station string, winds []Wind) {
	fmt.Println("Wind readings for station:", station)
	for _, w := range winds {
		fmt.Printf("%s\n", w.Time)
		fmt.Printf("\tSpeed/Gusts: %s/%s\n", w.Speed, w.Gusts)
		fmt.Printf("\tDirection: %s (%s)\n", w.DirectionDegrees, w.DirectionAcronym)
	}
}
