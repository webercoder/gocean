package waterlevel

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// WaterLevelsCommandHandler handles water levels commands.
type WaterLevelsCommandHandler struct {
	flagSet *flag.FlagSet
	predAPI *API
}

// NewCommandHandler creates a new Tides and Currents CommandHandler.
func NewCommandHandler() *WaterLevelsCommandHandler {
	return &WaterLevelsCommandHandler{
		flagSet: flag.NewFlagSet("tidesandcurrents", flag.ExitOnError),
		predAPI: NewAPI("gocean"),
	}
}

// GetFlagSet returns this command's flagSet for parsing command-line options.
func (wlch *WaterLevelsCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	return wlch.flagSet, nil
}

// HandleCommand .
func (wlch *WaterLevelsCommandHandler) HandleCommand(command string) error {
	station := wlch.flagSet.Arg(2)
	if len(station) == 0 {
		wlch.Usage(errors.New("station is required"))
		os.Exit(1)
	}

	results, err := wlch.predAPI.Retrieve(station, 24)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}
	wlch.predAPI.PrintTabDelimited(station, results)
	return nil
}

// Usage prints how to use this command.
func (wlch *WaterLevelsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
		fmt.Println("Usage:")
	}

	fmt.Println("waterlevels station-id")
}