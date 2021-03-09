package stations

import (
	"fmt"

	"github.com/webercoder/gocean/utils"
)

// CommandHandler .
type CommandHandler struct{}

// HandleCommand .
func (ch *CommandHandler) HandleCommand(arg string) error {
	fmt.Println("Finding nearest station to", arg)
	coords, err := utils.FindCoordsForPostcode(arg)
	if err != nil {
		return fmt.Errorf("Could not find coordinates for the provided location %s", arg)
	}

	stationsClient := NewClient()
	result := *stationsClient.GetNearestStation(*coords)
	fmt.Printf(
		"The nearest Station is \"%s\" (ID: %d), which is %f kms away from %s.\n",
		result.Station.Name,
		result.Station.ID,
		result.Distance,
		arg,
	)

	return nil
}
