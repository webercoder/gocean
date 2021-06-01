package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// AirTemperatureCommandHandler handles water levels commands.
type AirTemperatureCommandHandler struct {
	BaseCommandHandler
	AirTemperatureAPI *coops.AirTemperatureAPI
}

// NewAirTemperatureCommandHandler creates a new AirTemperatureCommandHandler.
func NewAirTemperatureCommandHandler() *AirTemperatureCommandHandler {
	return &AirTemperatureCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductAirTemperature),
		AirTemperatureAPI:  coops.NewAirTemperatureAPI("gocean"),
	}
}

// HandleCommand processes this command.
func (atch *AirTemperatureCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *AirTemperatureCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *AirTemperatureCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.AirTemperatureAPI.GetAirTemperatures(req)
	if err != nil {
		return err
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.AirTemperatureAPI.PrintTabDelimited(req, results)
	return nil
}

func (atch *AirTemperatureCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.AirTemperatureAPI.Client.Get(req)
	if err != nil {
		return err
	}

	fmt.Println(string(resp))
	return nil
}
