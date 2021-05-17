package lib

import (
	"fmt"
)

// FindCoordsForPostcode returns the geocoordinates for a zip code.
func FindCoordsForPostcode(zip string) (*GeoCoordinates, error) {
	if coords, ok := CensusZipCoords[zip]; ok {
		return &coords, nil
	}

	return nil, fmt.Errorf("Zip %s not found", zip)
}
