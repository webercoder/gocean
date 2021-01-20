package tides_test

import (
	"net/http"
	"testing"

	"github.com/webercoder/gocean/testutils"
	. "github.com/webercoder/gocean/tides"
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

type FakePredictionsClient struct {
	Err      error
	JsonData string
}

func (fsc *FakePredictionsClient) Get(url string) (resp *http.Response, err error) {
	if fsc.Err != nil {
		return nil, fsc.Err
	}

	return &http.Response{
		Body: testutils.NewStringReadCloser(fsc.JsonData),
	}, nil
}

func TestNOAATidesClient_RetrievePredictions(t *testing.T) {
	station := "9410170"
	client := &NOAATidesClient{Client: &FakePredictionsClient{JsonData: NOAAPredictionsJSONData}}
	data, err := client.RetrievePredictions(station, 1)
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
