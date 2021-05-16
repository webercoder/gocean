package lib_test

import (
	"math"
	"testing"

	"github.com/webercoder/gocean/lib"
)

func TestHarvestineDistance(t *testing.T) {
	sanDiegoCityHall := lib.GeoCoordinates{Lat: 32.716868, Long: -117.162837}
	lasVegasCityHall := lib.GeoCoordinates{Lat: 36.167206, Long: -115.148492}
	expected := 426.259297 // Expected from https://keisan.casio.com/exec/system/1224587128
	actual := lib.HarvesineDistance(sanDiegoCityHall, lasVegasCityHall)
	if math.Abs(actual-expected) > 0.001 {
		t.Errorf("The distance between San Diego and Las Vegas is %f not %f", expected, actual)
	}
}
