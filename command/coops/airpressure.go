package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// AirPressureCommandHandler handles air pressure commands.
type AirPressureCommandHandler struct {
	BaseCommandHandler
	AirPressureAPI *coops.AirPressureAPI
}

// NewAirPressureCommandHandler creates a new AirPressureCommandHandler.
func NewAirPressureCommandHandler() *AirPressureCommandHandler {
	return &AirPressureCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductAirPressure),
		AirPressureAPI:     coops.NewAirPressureAPI("gocean"),
	}
}

// HandleCommand .
func (atch *AirPressureCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *AirPressureCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *AirPressureCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.AirPressureAPI.GetAirPressure(req)
	if err != nil {
		return fmt.Errorf("could not load air pressure data for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.AirPressureAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *AirPressureCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.AirPressureAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve air pressure data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
