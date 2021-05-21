package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// WaterLevelsCommandHandler handles water levels commands.
type WaterLevelsCommandHandler struct {
	clientConfig  *CoopsClientConfig
	flagSet       *flag.FlagSet
	waterLevelAPI *coops.WaterLevelAPI
}

// NewWaterLevelsCommandHandler creates a new Tides and Currents CommandHandler.
func NewWaterLevelsCommandHandler() *WaterLevelsCommandHandler {
	clientConfig := NewCoopsClientConfig()
	return &WaterLevelsCommandHandler{
		clientConfig:  clientConfig,
		flagSet:       clientConfig.GetFlagSet("waterlevel", flag.ExitOnError),
		waterLevelAPI: coops.NewWaterLevelAPI("gocean"),
	}
}

// HandleCommand .
func (wlch *WaterLevelsCommandHandler) HandleCommand() error {
	if err := wlch.flagSet.Parse(os.Args[3:]); err != nil {
		wlch.Usage(errors.New("unable to parse command-line options"))
	}

	if wlch.clientConfig.Station == "" {
		wlch.Usage(errors.New("station is required"))
	}

	if wlch.clientConfig.BeginDate == "" && wlch.clientConfig.EndDate == "" {
		wlch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}

	reqOptions, err := wlch.clientConfig.ToRequestOptions()
	if err != nil {
		wlch.Usage(err)
	}
	req := coops.NewClientRequest(
		append(
			reqOptions,
			coops.WithFormat(coops.ResponseFormatJSON),
			coops.WithProduct(coops.ProductWaterLevel),
		)...,
	)

	results, err := wlch.waterLevelAPI.Retrieve(req)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}
	if wlch.clientConfig.Count > 0 {
		results = results[:wlch.clientConfig.Count]
	}
	wlch.waterLevelAPI.PrintTabDelimited(wlch.clientConfig.Station, results)
	return nil
}

// Usage prints how to use this command.
func (wlch *WaterLevelsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	wlch.flagSet.Usage()
	os.Exit(1)
}
