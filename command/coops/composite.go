package coops

import (
	"errors"
	"fmt"
	"os"

	"github.com/webercoder/gocean/command"
	"github.com/webercoder/gocean/src/coops"
)

// CompositeCommandHandler is a composite of all the CO-OPS command-line commands.
type CompositeCommandHandler struct {
	subHandlers map[coops.Product]command.CoopsHandler
}

// NewCompositeCommandHandler creates a new CompositeCommandHandler.
func NewCompositeCommandHandler() *CompositeCommandHandler {
	return &CompositeCommandHandler{
		subHandlers: map[coops.Product]command.CoopsHandler{
			coops.ProductAirGap:           NewAirGapCommandHandler(),
			coops.ProductAirPressure:      NewAirPressureCommandHandler(),
			coops.ProductAirTemperature:   NewAirTemperatureCommandHandler(),
			coops.ProductConductivity:     NewConductivityCommandHandler(),
			coops.ProductPredictions:      NewPredictionsCommandHandler(),
			coops.ProductWaterLevel:       NewWaterLevelsCommandHandler(),
			coops.ProductWaterTemperature: NewWaterTemperatureCommandHandler(),
			coops.ProductWind:             NewWindCommandHandler(),
		},
	}
}

// GetRequestOptions is just to satisfy the interface, it should only be called by the subhandler.
func (cch *CompositeCommandHandler) GetRequestOptions() []coops.ClientRequestOption {
	panic(errors.New("should not call CompositeCommandHandler.GetRequestOptions directly"))
}

// HandleCommand calls the appropriate subcommand.
func (cch *CompositeCommandHandler) HandleCommand() error {
	if len(os.Args) < 3 {
		cch.Usage(errors.New("please provide a subcommand"))
	}
	command, ok := coops.StringToProduct(os.Args[2])
	if !ok {
		cch.Usage(fmt.Errorf("unknown product: %s", command))
	}

	handler, ok := cch.subHandlers[command]
	if !ok {
		cch.Usage(fmt.Errorf("subcommand %s is not supported", command))
	}

	return handler.HandleCommand()
}

// ParseFlags is just to satisfy the interface, it should only be called by the subhandler.
func (cch *CompositeCommandHandler) ParseFlags() {
	panic(errors.New("should not call CompositeCommandHandler.ParseFlags directly"))
}

// Usage prints the usage for all subcommands of this command.
func (cch *CompositeCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	for k := range cch.subHandlers {
		fmt.Printf("  %s\n", k)
	}

	os.Exit(1)
}
