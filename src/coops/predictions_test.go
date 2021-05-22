package coops_test

import (
	"testing"

	"github.com/webercoder/gocean/src/coops"
)

const NOAAPredictionsJSONData = `{
    "predictions": [{
        "t": "2021-01-19 07:00",
        "v": "1.894"
    }, {
        "t": "2021-01-19 07:06",
        "v": "1.888"
    }, {
        "t": "2021-01-19 07:12",
        "v": "1.882"
    }, {
        "t": "2021-01-19 07:18",
        "v": "1.877"
    }, {
        "t": "2021-01-19 07:24",
        "v": "1.873"
    }]
}`

func TestPredictionsRetrieve(t *testing.T) {
	api := &coops.PredictionsAPI{
		App:    "gocean_test",
		Client: &coops.Client{HTTPClient: &FakeCoopsClient{JsonData: NOAAPredictionsJSONData}},
	}
	station := "9410170"
	data, err := api.GetPredictions(coops.NewClientRequest(
		coops.WithStation(station),
		coops.WithHours(1),
	))
	if err != nil {
		t.Error("Did not expect error when retrieving tide data", err)
	}
	if len(data) != 5 {
		t.Error("Zero tides received", data)
	}
	if data[2].Time != "2021-01-19 07:12" {
		t.Error("Unexpected time for the third entry")
	}
	if data[2].Value != "1.882" {
		t.Error("Unexpected value for the third entry")
	}
}
