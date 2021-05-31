package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAAirGapJSONData = `{
    "metadata": {
        "id": "9410230",
        "name": "La Jolla",
        "lat": "32.8669",
        "lon": "-117.2571"
    },
    "data": [{
        "t": "2021-05-23 06:48",
        "v": "137.936",
        "s": "0.020",
        "f": "0,0,0,0"
    }, {
        "t": "2021-05-23 06:54",
        "v": "137.792",
        "s": "0.030",
        "f": "0,0,0,0"
    }, {
        "t": "2021-05-23 07:00",
        "v": "137.612",
        "s": "0.036",
        "f": "0,0,0,0"
    }, {
        "t": "2021-05-23 07:06",
        "v": "137.323",
        "s": "0.052",
        "f": "0,0,0,0"
    }, {
        "t": "2021-05-23 07:12",
        "v": "137.064",
        "s": "0.043",
        "f": "0,0,0,0"
    }]
}`

func TestGetAirGaps(t *testing.T) {
	api := &coops.AirGapAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAAirGapJSONData}},
	}
	station := "9410170"
	data, err := api.GetAirGap(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving air gap data", err)
	}
	if len(data) != 5 {
		t.Error("Incorrect number of air gap received", data)
	}
	if data[2].Time != "2021-05-23 07:00" {
		t.Error("Unexpected time for the third entry", data[2].Time)
	}
	if data[2].Value != "137.612" {
		t.Error("Unexpected value for the third entry", data[2].Value)
	}
	if data[2].Sigma != "0.036" {
		t.Error("Unexpected sigma for the third entry", data[2].Sigma)
	}
	if data[2].Flags != "0,0,0,0" {
		t.Error("Unexpected flags for the third entry", data[2].Flags)
	}
}
