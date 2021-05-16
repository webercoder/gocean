package predictions

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// CommandHandler .
type PredictionsCommandHandler struct {
	flagSet *flag.FlagSet
	predApi *PredictionsApi
}

// NewCommandHandler creates a new Tides and Currents CommandHandler
func NewCommandHandler() *PredictionsCommandHandler {
	return &PredictionsCommandHandler{
		flagSet: flag.NewFlagSet("tidesandcurrents", flag.ExitOnError),
		predApi: NewPredictionApi("gocean"),
	}
}

// GetFlagSet returns this command's flagSet for parsing command line options.
func (ch *PredictionsCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	return ch.flagSet, nil
}

// HandleCommand .
func (pch *PredictionsCommandHandler) HandleCommand(command string) error {
	station := pch.flagSet.Arg(2)
	if len(station) == 0 {
		pch.Usage(errors.New("station is required"))
		os.Exit(1)
	}

	results, err := pch.predApi.Retrieve(station, 24)
	if err != nil {
		return fmt.Errorf("Could not load predictions for station")
	}
	pch.predApi.PrintTabDelimited(station, results)
	return nil
}

func (ch *PredictionsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
		fmt.Println("Usage:")
	}

	fmt.Println("predictions station-id")
}
