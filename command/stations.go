package command

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/src/lib"
	"github.com/webercoder/gocean/src/stations"
)

// StationsCommandHandler handles stations commands.
type StationsCommandHandler struct {
	flagSet  *flag.FlagSet
	postcode string
}

// NewStationsCommandHandler creates a new Stations CommandHandler.
func NewStationsCommandHandler() *StationsCommandHandler {
	handler := &StationsCommandHandler{
		flagSet: flag.NewFlagSet("stations", flag.ExitOnError),
	}
	handler.flagSet.StringVar(&handler.postcode, "postcode", "", "Find stations near this postcode")
	return handler
}

// HandleCommand handles the stations command.
func (sch *StationsCommandHandler) HandleCommand() error {
	if err := sch.flagSet.Parse(os.Args[2:]); err != nil {
		sch.Usage(errors.New("unable to parse command-line options"))
	}

	if len(sch.postcode) == 0 {
		sch.Usage(errors.New("postcode is required"))
		os.Exit(1)
	}

	fmt.Println("Finding nearest station to", sch.postcode)

	coords, err := lib.FindCoordsForPostcode(sch.postcode)
	if err != nil {
		return fmt.Errorf("could not find coordinates for the provided location %s", sch.postcode)
	}

	stationsClient := stations.NewClient()
	result := *stationsClient.GetNearestStation(*coords)
	fmt.Printf(
		"The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n",
		result.Station.Name,
		result.Station.ID,
		result.Distance,
		sch.postcode,
	)

	return nil
}

// Usage prints how to use this command.
func (sch *StationsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	sch.flagSet.Usage()
	os.Exit(1)
}
