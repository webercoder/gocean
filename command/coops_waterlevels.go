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
	wlch.validate()

	req := coops.NewClientRequest(wlch.getRequestOptions()...)

	if wlch.clientConfig.Format == ResponseFormatPrettyPrint {
		return wlch.handlePrettyPrint(req)
	}

	return wlch.handleRawPrint(req)
}

func (wlch *WaterLevelsCommandHandler) validate() {
	if err := wlch.flagSet.Parse(os.Args[3:]); err != nil {
		wlch.Usage(errors.New("unable to parse command-line options"))
	}

	if wlch.clientConfig.Station == "" {
		wlch.Usage(errors.New("station is required"))
	}

	if wlch.clientConfig.BeginDate == "" && wlch.clientConfig.EndDate == "" {
		wlch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (wlch *WaterLevelsCommandHandler) getRequestOptions() []coops.ClientRequestOption {
	reqOptions, err := wlch.clientConfig.ToRequestOptions()
	if err != nil {
		wlch.Usage(err)
	}

	return append(reqOptions, coops.WithProduct(coops.ProductWaterLevel))
}

func (wlch *WaterLevelsCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := wlch.waterLevelAPI.GetWaterLevels(req)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}

	if wlch.clientConfig.Count > 0 {
		results = results[:wlch.clientConfig.Count]
	}

	wlch.waterLevelAPI.PrintTabDelimited(wlch.clientConfig.Station, results)
	return nil
}

func (wlch *WaterLevelsCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := wlch.waterLevelAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve water_level data: %v", err)
	}

	fmt.Println(string(resp))
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
