package lib_test

import (
	"testing"

	"github.com/webercoder/gocean/lib"
)

func TestCensusZipCoords(t *testing.T) {
	tests := map[string]lib.GeoCoordinates{
		"92101": {Lat: 32.724103, Long: -117.170912},
		"98101": {Lat: 47.610902, Long: -122.336422},
		"18837": {Lat: 41.915072, Long: -76.30286},
	}

	for zip, expected := range tests {
		actual, err := lib.FindCoordsForPostcode(zip)
		if err != nil {
			t.Errorf("Error retrieving coords for zip code %s: %s", zip, err)
		}

		if actual.Lat != expected.Lat || actual.Long != expected.Long {
			t.Errorf("Expected %+v to equal %+v", actual, expected)
		}
	}
}
