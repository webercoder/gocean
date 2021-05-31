package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAVisibilityJSONData = `{
    "metadata": {
        "id": "9410230",
        "name": "La Jolla",
        "lat": "32.8669",
        "lon": "-117.2571"
    },
    "data": [{
        "t": "2021-05-23 06:48",
        "v": "5.40",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 06:54",
        "v": "5.40",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:00",
        "v": "5.40",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:06",
        "v": "5.40",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:12",
        "v": "5.40",
        "f": "0,0,0"
    }]
}`

func TestGetVisibility(t *testing.T) {
	api := &coops.VisibilityAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAVisibilityJSONData}},
	}
	station := "9410170"
	data, err := api.GetVisibility(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving air visibility data", err)
	}
	if len(data) != 5 {
		t.Error("Incorrect number of air visibility received", data)
	}
	if data[2].Time != "2021-05-23 07:00" {
		t.Error("Unexpected time for the third entry", data[2].Time)
	}
	if data[2].Value != "5.40" {
		t.Error("Unexpected value for the third entry", data[2].Value)
	}
	if data[2].Flags != "0,0,0" {
		t.Error("Unexpected flags for the third entry", data[2].Flags)
	}
}
