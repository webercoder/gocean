package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAWindJSONData = `{
    "metadata": {
        "id": "9410230",
        "name": "La Jolla",
        "lat": "32.8669",
        "lon": "-117.2571"
    },
    "data": [{
        "t": "2021-05-23 06:48",
        "s": "1.36",
        "d": "163.00",
        "dr": "SSE",
        "g": "1.94",
        "f": "0,0"
    }, {
        "t": "2021-05-23 06:54",
        "s": "0.78",
        "d": "84.00",
        "dr": "E",
        "g": "1.94",
        "f": "0,0"
    }, {
        "t": "2021-05-23 07:00",
        "s": "1.75",
        "d": "18.00",
        "dr": "NNE",
        "g": "3.89",
        "f": "0,0"
    }, {
        "t": "2021-05-23 07:06",
        "s": "1.94",
        "d": "28.00",
        "dr": "NNE",
        "g": "3.11",
        "f": "0,0"
    }, {
        "t": "2021-05-23 07:12",
        "s": "2.53",
        "d": "34.00",
        "dr": "NE",
        "g": "3.30",
        "f": "0,0"
    }]
}`

func TestGetWinds(t *testing.T) {
	api := &coops.WindAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAWindJSONData}},
	}
	station := "9410170"
	data, err := api.GetWind(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving wind wind data", err)
	}
	if len(data) != 5 {
		t.Error("Incorrect number of winds received", data)
	}
	if data[2].Time != "2021-05-23 07:00" {
		t.Error("Unexpected time for the third entry", data[2].Time)
	}
	if data[2].Speed != "1.75" {
		t.Error("Unexpected speed for the third entry", data[2].Speed)
	}
	if data[2].DirectionDegrees != "18.00" {
		t.Error("Unexpected direction degrees for the third entry", data[2].Speed)
	}
	if data[2].DirectionAcronym != "NNE" {
		t.Error("Unexpected direction acronym for the third entry", data[2].Speed)
	}
	if data[2].Gusts != "3.89" {
		t.Error("Unexpected gusts for the third entry", data[2].Speed)
	}
	if data[2].Flags != "0,0" {
		t.Error("Unexpected f for the third entry", data[2].Flags)
	}
}
