package utils

import (
	"math"
)

// GeoCoordinates represent a latitude and longitude on Earth.
type GeoCoordinates struct {
	Lat  float64
	Long float64
}

// HarvesineDistance finds the distance in KMs between two points on Earth.
// Formula converted to golang from https://www.movable-type.co.uk/scripts/latlong.html
func HarvesineDistance(coords1, coords2 GeoCoordinates) float64 { //lat1, long1, lat2, long2 float64) float64 {
	earthsRadiusKM := 6378.14
	φ1 := coords1.Lat * math.Pi / 180 // φ, λ in radians
	φ2 := coords2.Lat * math.Pi / 180
	Δφ := (coords2.Lat - coords1.Lat) * math.Pi / 180
	Δλ := (coords2.Long - coords1.Long) * math.Pi / 180
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))
	return earthsRadiusKM * c
}
