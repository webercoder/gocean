package main

import (
	"fmt"
	"os"

	"github.com/webercoder/gocean/stations"
	"github.com/webercoder/gocean/utils"
)

func usage(msg string, code int) {
	if len(msg) > 0 {
		fmt.Printf("%s\n", msg)
	}
	fmt.Printf("Usage: %s zipcode", os.Args[0])
	os.Exit(1)
}

func main() {
	// temporary, replace with proper command line parsing
	args := os.Args[1:]
	if len(args) < 1 {
		usage("", 1)
	}

	zip := os.Args[1]
	coords, err := utils.FindCoordsForZip(zip)
	if err != nil {
		usage("Could not find coordinates for the provided zip", 2)
	}

	stationManager := stations.NewNOAAStationManager()
	station, distance := stationManager.GetNearestStation(*coords)
	fmt.Printf("The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n", station.Name, station.ID, distance, zip)
}
