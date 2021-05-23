package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAAirTemperatureJSONData = `{
    "metadata": {
        "id": "9410170",
        "name": "La Jolla",
        "lat": "32.8669",
        "lon": "-117.2571"
    },
    "data": [{
        "t": "2020-05-21 06:48",
        "v": "60.8",
        "f": "0,0,0"
    }, {
        "t": "2020-05-21 06:54",
        "v": "60.6",
        "f": "0,0,0"
    }, {
        "t": "2020-05-21 07:00",
        "v": "61.5",
        "f": "0,0,0"
    }, {
        "t": "2020-05-21 07:06",
        "v": "61.7",
        "f": "0,0,0"
    }, {
        "t": "2020-05-21 07:12",
        "v": "61.5",
        "f": "0,0,0"
    }]
}`

func TestGetAirTemperatures(t *testing.T) {
	api := &coops.AirTemperatureAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAAirTemperatureJSONData}},
	}
	station := "9410170"
	data, err := api.GetAirTemperatures(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving air temp data", err)
	}
	if len(data) != 5 {
		t.Error("Zero temps received", data)
	}
	if data[2].Time != "2020-05-21 07:00" {
		t.Error("Unexpected time for the third entry", data[2].Time)
	}
	if data[2].Value != "61.5" {
		t.Error("Unexpected value for the third entry", data[2].Value)
	}
}
