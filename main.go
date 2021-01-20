package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/noaa/stations"
	"github.com/webercoder/gocean/noaa/tidesandcurrents"
	"github.com/webercoder/gocean/noaa/tidesandcurrents/waterlevels"
	"github.com/webercoder/gocean/utils"
)

func usage(msg string) {
	if len(msg) > 0 {
		fmt.Printf("%s\n", msg)
	}
	fmt.Printf("Usage:\n\t%s stations postcode\n\t%s tides station-id\n", os.Args[0], os.Args[0])
	os.Exit(1)
}

func handleTidesCommand(station string) {
	tidesClient := tidesandcurrents.NewClient()
	predictions, err := waterlevels.RetrievePredictions(tidesClient, station, 24)
	if err != nil {
		usage("Could not load predictions for station")
	}
	waterlevels.PrintPredictions(station, predictions)
}

func handleStationCommand(location string) {
	fmt.Println("Finding nearest station to", location)
	coords, err := utils.FindCoordsForPostcode(location)
	if err != nil {
		usage("Could not find coordinates for the provided location")
	}

	stationsClient := stations.NewClient()
	result := *stationsClient.GetNearestStation(*coords)
	fmt.Printf(
		"The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n",
		result.Station.Name,
		result.Station.ID,
		result.Distance,
		location,
	)
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
	stationCmd := flag.NewFlagSet("stations", flag.ExitOnError)
	tideCmd := flag.NewFlagSet("tides", flag.ExitOnError)
	configCmd := flag.NewFlagSet("config", flag.ExitOnError)

	if len(os.Args) < 2 {
		usage("")
	}

	switch os.Args[1] {
	case "stations":
		if err := stationCmd.Parse(os.Args[2:]); err != nil {
			usage("Unable to parse stations commands")
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
