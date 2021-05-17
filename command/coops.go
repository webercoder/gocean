package command

import (
	"errors"
	"flag"
	"fmt"
)

// COOPSCommandHandler is a composite of all the CO-OPS command-line commands.
type COOPSCommandHandler struct {
	subHandlers map[string]Handler
}

// NewCOOPSCommandHandler creates a new Tides and Currents CommandHandler.
func NewCOOPSCommandHandler() *COOPSCommandHandler {
	return &COOPSCommandHandler{
		subHandlers: map[string]Handler{
			"predictions": NewPredictionsCommandHandler(),
			"waterlevels": NewWaterLevelsCommandHandler(),
		},
	}
}

// GetFlagSet returns the flag set for the subcommand.
func (cch *COOPSCommandHandler) GetFlagSet(command string) (*flag.FlagSet, error) {
	if command == "" {
		return nil, errors.New("please provide a subcommand")
	}

	handler, ok := cch.subHandlers[command]
	if !ok {
		return nil, fmt.Errorf("command %s is not supported", command)
	}

	return handler.GetFlagSet("")
}

// HandleCommand calls the appropriate subcommand.
func (cch *COOPSCommandHandler) HandleCommand(command string) error {
	if command == "" {
		return errors.New("please provide a subcommand")
	}

	handler, ok := cch.subHandlers[command]
	if !ok {
		return fmt.Errorf("subcommand %s is not supported", command)
	}

	return handler.HandleCommand("")
}

// Usage prints the usage for all subcommands of this command.
func (cch *COOPSCommandHandler) Usage(err ...error) {
	for key := range cch.subHandlers {
		fmt.Printf("tidesandcurrents %s ...\n", key)
	}
}
