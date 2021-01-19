package tides_test

import (
	"testing"

	. "github.com/webercoder/gocean/tides"
)

func TestNOAATidesClient_RetrievePredictions(t *testing.T) {
	station := "9410170"
	client := NewNOAATidesClient()
	data, err := client.RetrievePredictions(station, 1)
	if err != nil {
		t.Error("Did not expect error when retrieving tide data", err)
	}
	if len(data) == 0 {
		t.Error("Zero tides received", data)
	}
}
