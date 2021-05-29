package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// WaterTemperatureCommandHandler handles water levels commands.
type WaterTemperatureCommandHandler struct {
	clientConfig        *CoopsClientConfig
	flagSet             *flag.FlagSet
	WaterTemperatureAPI *coops.WaterTemperatureAPI
}

// NewWaterTemperatureCommandHandler creates a new Tides and Currents CommandHandler.
func NewWaterTemperatureCommandHandler() *WaterTemperatureCommandHandler {
	clientConfig := NewCoopsClientConfig()
	return &WaterTemperatureCommandHandler{
		clientConfig:        clientConfig,
		flagSet:             clientConfig.GetFlagSet(coops.ProductWaterTemperature.String(), flag.ExitOnError),
		WaterTemperatureAPI: coops.NewWaterTemperatureAPI("gocean"),
	}
}

// HandleCommand .
func (atch *WaterTemperatureCommandHandler) HandleCommand() error {
	atch.validate()

	req := coops.NewClientRequest(atch.getRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *WaterTemperatureCommandHandler) validate() {
	if err := atch.flagSet.Parse(os.Args[3:]); err != nil {
		atch.Usage(errors.New("unable to parse command-line options"))
	}

	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *WaterTemperatureCommandHandler) getRequestOptions() []coops.ClientRequestOption {
	reqOptions, err := atch.clientConfig.ToRequestOptions()
	if err != nil {
		atch.Usage(err)
	}

	return append(reqOptions, coops.WithProduct(coops.ProductWaterTemperature))
}

func (atch *WaterTemperatureCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.WaterTemperatureAPI.GetWaterTemperatures(req)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.WaterTemperatureAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *WaterTemperatureCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.WaterTemperatureAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve water_level data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}

// Usage prints how to use this command.
func (atch *WaterTemperatureCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	atch.flagSet.Usage()
	os.Exit(1)
}
