package predictions_test

import (
	"net/http"
	"testing"

	"github.com/webercoder/gocean/lib"
	tclib "github.com/webercoder/gocean/noaa/tidesandcurrents/lib"
	"github.com/webercoder/gocean/noaa/tidesandcurrents/predictions"
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

type FakeTidesAndCurrentsClient struct {
	Err      error
	JsonData string
}

func (fsc *FakeTidesAndCurrentsClient) Get(url string) (resp *http.Response, err error) {
	if fsc.Err != nil {
		return nil, fsc.Err
	}

	return &http.Response{
		Body: lib.NewStringReadCloser(fsc.JsonData),
	}, nil
}

func TestRetrieve(t *testing.T) {
	api := &predictions.PredictionsApi{
		App:    "gocean_test",
		Client: &tclib.Client{HTTPClient: &FakeTidesAndCurrentsClient{JsonData: NOAAPredictionsJSONData}},
	}
	station := "9410170"
	data, err := api.Retrieve(station, 1)
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
