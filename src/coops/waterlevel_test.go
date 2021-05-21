package coops_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAWaterLevelsJSONData = `{
    "metadata": {
        "id": "9410170",
        "name": "San Diego, San Diego Bay",
        "lat": "32.7142",
        "lon": "-117.1736"
    }, 
    "data": [{
        "t": "2021-01-19 07:00",
        "v": "1.894",
        "q": "v"
    }, {
        "t": "2021-01-19 07:06",
        "v": "1.888",
        "q": "v"
    }, {
        "t": "2021-01-19 07:12",
        "v": "1.882",
        "q": "p"
    }, {
        "t": "2021-01-19 07:18",
        "v": "1.877",
        "q": "p"
    }, {
        "t": "2021-01-19 07:24",
        "v": "1.873",
        "q": "p"
    }]
}`

const NOAAWaterLevelsJSONErrorData = `{
	"error": {
		"message": "No data was found. This product may not be offered at this station at the requested time."
	}
}`

func TestRetrieve(t *testing.T) {
	api := &coops.WaterLevelAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAWaterLevelsJSONData}},
	}
	station := "9410170"
	data, err := api.Retrieve(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving tide data", err)
	}
	if len(data) != 5 {
		t.Error("Zero tides received", data)
	}
	if data[1].Quality != "v" {
		t.Error("Unexpected quality for the second entry")
	}
	if data[2].Time != "2021-01-19 07:12" {
		t.Error("Unexpected time for the third entry")
	}
	if data[2].Value != "1.882" {
		t.Error("Unexpected value for the third entry")
	}
	if data[2].Quality != "p" {
		t.Error("Unexpected quality for the third entry")
	}
}

func TestRetrieveError(t *testing.T) {
	api := &coops.WaterLevelAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAWaterLevelsJSONErrorData}},
	}
	station := "9410170"
	_, err := api.Retrieve(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err == nil {
		t.Error("Expected an error when retrieving bad water level data")
	}
	if !strings.Contains(fmt.Sprint(err), "No data was found") {
		t.Error("Error should contain the correct message")
	}
}
