package predictions

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// CommanderHandler handles predictions commands.
type CommanderHandler struct {
	flagSet *flag.FlagSet
	predAPI *API
}

// NewCommandHandler creates a new Tides and Currents CommandHandler
func NewCommandHandler() *CommanderHandler {
	return &CommanderHandler{
		flagSet: flag.NewFlagSet("tidesandcurrents", flag.ExitOnError),
		predAPI: NewAPI("gocean"),
	}
}

// GetFlagSet returns this command's flagSet for parsing command-line options.
func (pch *CommanderHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	return pch.flagSet, nil
}

// HandleCommand processes the predictions command.
func (pch *CommanderHandler) HandleCommand(command string) error {
	station := pch.flagSet.Arg(2)
	if len(station) == 0 {
		pch.Usage(errors.New("station is required"))
		os.Exit(1)
	}

	results, err := pch.predAPI.Retrieve(station, 24)
	if err != nil {
		return fmt.Errorf("could not load predictions for station")
	}
	pch.predAPI.PrintTabDelimited(station, results)
	return nil
}

// Usage prints how to use this command.
func (pch *CommanderHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
		fmt.Println("Usage:")
	}

	fmt.Println("predictions station-id")
}
