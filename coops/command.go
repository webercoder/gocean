package coops

import (
	"errors"
	"flag"
	"fmt"

	"github.com/webercoder/gocean/coops/predictions"
	"github.com/webercoder/gocean/coops/water_level"
	"github.com/webercoder/gocean/lib"
)

// CommandHandler .
type TidesAndCurrentsCommandHandler struct {
	subHandlers map[string]lib.CommandHandler
}

// NewCommandHandler creates a new Tides and Currents CommandHandler
func NewTidesAndCurrentsCommandHandler() *TidesAndCurrentsCommandHandler {
	return &TidesAndCurrentsCommandHandler{
		subHandlers: map[string]lib.CommandHandler{
			"predictions": predictions.NewCommandHandler(),
			"waterlevels": water_level.NewCommandHandler(),
		},
	}
}

func (tch *TidesAndCurrentsCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	if command == "" {
		return nil, errors.New("please provide a subcommand")
	}

	handler, ok := tch.subHandlers[command]
	if !ok {
		return nil, fmt.Errorf("command %s is not supported", command)
	}

	return handler.GetFlagSet("")
}

// HandleCommand .
func (tch *TidesAndCurrentsCommandHandler) HandleCommand(command string) error {
	if command == "" {
		return errors.New("please provide a subcommand")
	}

	handler, ok := tch.subHandlers[command]
	if !ok {
		return fmt.Errorf("subcommand %s is not supported", command)
	}

	return handler.HandleCommand("")
}

func (tch *TidesAndCurrentsCommandHandler) Usage(err ...error) {
	for key := range tch.subHandlers {
		fmt.Printf("tidesandcurrents %s ...\n", key)
	}
}
