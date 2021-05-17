package coops

import (
	"errors"
	"flag"
	"fmt"

	"github.com/webercoder/gocean/coops/predictions"
	"github.com/webercoder/gocean/coops/waterlevel"
	"github.com/webercoder/gocean/lib"
)

// TidesAndCurrentsCommandHandler is a composite of all the CO-OPS command-line commands.
type TidesAndCurrentsCommandHandler struct {
	subHandlers map[string]lib.CommandHandler
}

// NewTidesAndCurrentsCommandHandler creates a new Tides and Currents CommandHandler.
func NewTidesAndCurrentsCommandHandler() *TidesAndCurrentsCommandHandler {
	return &TidesAndCurrentsCommandHandler{
		subHandlers: map[string]lib.CommandHandler{
			"predictions": predictions.NewCommandHandler(),
			"waterlevels": waterlevel.NewCommandHandler(),
		},
	}
}

// GetFlagSet returns the flag set for the subcommand.
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

// HandleCommand calls the appropriate subcommand.
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

// Usage prints the usage for all subcommands of this command.
func (tch *TidesAndCurrentsCommandHandler) Usage(err ...error) {
	for key := range tch.subHandlers {
		fmt.Printf("tidesandcurrents %s ...\n", key)
	}
}
