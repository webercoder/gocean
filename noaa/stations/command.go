package stations

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/utils"
)

// CommandHandler .
type StationsCommandHandler struct {
	flagSet *flag.FlagSet
}

// NewCommandHandler creates a new Stations CommandHandler.
func NewStationsCommandHandler() *StationsCommandHandler {
	return &StationsCommandHandler{
		flagSet: flag.NewFlagSet("stations", flag.ExitOnError),
	}
}

// GetFlagSet returns this command's flagSet for parsing command line options.
func (sch *StationsCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	return sch.flagSet, nil
}

// HandleCommand handles the stations command.
func (sch *StationsCommandHandler) HandleCommand(command string) error {
	postcode := sch.flagSet.Arg(1)
	if len(postcode) == 0 {
		sch.Usage(errors.New("postcode is required"))
		os.Exit(1)
	}

	fmt.Println("Finding nearest station to", postcode)

	coords, err := utils.FindCoordsForPostcode(postcode)
	if err != nil {
		return fmt.Errorf("Could not find coordinates for the provided location %s", postcode)
	}

	stationsClient := NewClient()
	result := *stationsClient.GetNearestStation(*coords)
	fmt.Printf(
		"The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n",
		result.Station.Name,
		result.Station.ID,
		result.Distance,
		postcode,
	)

	return nil
}

func (sch *StationsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
		fmt.Println("Usage:")
	}

	fmt.Println("stations postcode")
}
