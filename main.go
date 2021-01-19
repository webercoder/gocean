package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/stations"
	"github.com/webercoder/gocean/tides"
	"github.com/webercoder/gocean/utils"
)

func usage(msg string) {
	if len(msg) > 0 {
		fmt.Printf("%s\n", msg)
	}
	fmt.Printf("Usage:\n\t%s station postcode\n\t%s tides station-id\n", os.Args[0], os.Args[0])
	os.Exit(1)
}

func handleTidesCommand(station string) {
	tidesClient := tides.NewNOAATidesClient()
	predictions, err := tidesClient.RetrievePredictions(station, 48)
	if err != nil {
		usage("Could not load predictions for station")
	}
	tides.PrintPredictions(station, predictions)
}

func handleStationCommand(location string) {
	fmt.Println("Finding nearest station to", location)
	coords, err := utils.FindCoordsForPostcode(location)
	if err != nil {
		usage("Could not find coordinates for the provided location")
	}

	stationManager := stations.NewNOAAStationClient()
	station, distance := stationManager.GetNearestStation(*coords)
	fmt.Printf("The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n", station.Name, station.ID, distance, location)
}

func config(command string) {
	switch command {
	case "init":
	case "reload-geo-data":
	case "reload-station-data":
	default:
		usage("Unknown subcommand")
	}
}

func parseArgs() {
	stationCmd := flag.NewFlagSet("station", flag.ExitOnError)
	tideCmd := flag.NewFlagSet("tides", flag.ExitOnError)
	configCmd := flag.NewFlagSet("config", flag.ExitOnError)

	if len(os.Args) < 2 {
		usage("")
	}

	switch os.Args[1] {
	case "station":
		if err := stationCmd.Parse(os.Args[2:]); err != nil {
			usage("Unable to parse station commands")
		}
		handleStationCommand(stationCmd.Arg(0))
	case "tides":
		if err := tideCmd.Parse(os.Args[2:]); err != nil {
			usage("Unable to parse tides commands")
		}
		handleTidesCommand(tideCmd.Arg(0))
	case "config":
		if err := configCmd.Parse(os.Args[2:]); err != nil {
			usage("Unable to parse config commands")
		}
		usage("Config command is not currently supported")
	default:
		usage("No commands found")
	}
}

func main() {
	parseArgs()
}
