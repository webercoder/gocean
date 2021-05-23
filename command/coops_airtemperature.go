package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// AirTemperatureCommandHandler handles water levels commands.
type AirTemperatureCommandHandler struct {
	clientConfig      *CoopsClientConfig
	flagSet           *flag.FlagSet
	AirTemperatureAPI *coops.AirTemperatureAPI
}

// NewAirTemperatureCommandHandler creates a new Tides and Currents CommandHandler.
func NewAirTemperatureCommandHandler() *AirTemperatureCommandHandler {
	clientConfig := NewCoopsClientConfig()
	return &AirTemperatureCommandHandler{
		clientConfig:      clientConfig,
		flagSet:           clientConfig.GetFlagSet(coops.ProductAirTemperature.String(), flag.ExitOnError),
		AirTemperatureAPI: coops.NewAirTemperatureAPI("gocean"),
	}
}

// HandleCommand .
func (atch *AirTemperatureCommandHandler) HandleCommand() error {
	atch.validate()

	req := coops.NewClientRequest(atch.getRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *AirTemperatureCommandHandler) validate() {
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

func (atch *AirTemperatureCommandHandler) getRequestOptions() []coops.ClientRequestOption {
	reqOptions, err := atch.clientConfig.ToRequestOptions()
	if err != nil {
		atch.Usage(err)
	}

	return append(reqOptions, coops.WithProduct(coops.ProductAirTemperature))
}

func (atch *AirTemperatureCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.AirTemperatureAPI.GetAirTemperatures(req)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.AirTemperatureAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *AirTemperatureCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.AirTemperatureAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve water_level data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}

// Usage prints how to use this command.
func (atch *AirTemperatureCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	atch.flagSet.Usage()
	os.Exit(1)
}
