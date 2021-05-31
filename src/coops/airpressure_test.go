package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAAirPressureJSONData = `{
    "metadata": {
        "id": "9410230",
        "name": "La Jolla",
        "lat": "32.8669",
        "lon": "-117.2571"
    },
    "data": [{
        "t": "2021-05-23 06:48",
        "v": "1019.0",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 06:54",
        "v": "1019.0",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:00",
        "v": "1019.1",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:06",
        "v": "1019.2",
        "f": "0,0,0"
    }, {
        "t": "2021-05-23 07:12",
        "v": "1019.2",
        "f": "0,0,0"
    }]
}`

func TestGetAirPressures(t *testing.T) {
	api := &coops.AirPressureAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAAirPressureJSONData}},
	}
	station := "9410170"
	data, err := api.GetAirPressure(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving air pressure data", err)
	}
	if len(data) != 5 {
		t.Error("Incorrect number of air pressure received", data)
	}
	if data[2].Time != "2021-05-23 07:00" {
		t.Error("Unexpected time for the third entry", data[2].Time)
	}
	if data[2].Value != "1019.1" {
		t.Error("Unexpected value for the third entry", data[2].Value)
	}
}
