package water_level

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// CommandHandler .
type WaterLevelsCommandHandler struct {
	flagSet *flag.FlagSet
	predAPI *WaterLevelAPI
}

// NewCommandHandler creates a new Tides and Currents CommandHandler
func NewCommandHandler() *WaterLevelsCommandHandler {
	return &WaterLevelsCommandHandler{
		flagSet: flag.NewFlagSet("tidesandcurrents", flag.ExitOnError),
		predAPI: NewWaterLevelAPI("gocean"),
	}
}

// GetFlagSet returns this command's flagSet for parsing command line options.
func (ch *WaterLevelsCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	return ch.flagSet, nil
}

// HandleCommand .
func (pch *WaterLevelsCommandHandler) HandleCommand(command string) error {
	station := pch.flagSet.Arg(2)
	if len(station) == 0 {
		pch.Usage(errors.New("station is required"))
		os.Exit(1)
	}

	results, err := pch.predAPI.Retrieve(station, 24)
	if err != nil {
		return fmt.Errorf("Could not load water levels for station: %v", err)
	}
	pch.predAPI.PrintTabDelimited(station, results)
	return nil
}

func (ch *WaterLevelsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
		fmt.Println("Usage:")
	}

	fmt.Println("waterlevels station-id")
}
