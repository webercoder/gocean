package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// AirGapCommandHandler handles air gap commands.
type AirGapCommandHandler struct {
	BaseCommandHandler
	AirGapAPI *coops.AirGapAPI
}

// NewAirGapCommandHandler creates a new AirGapCommandHandler.
func NewAirGapCommandHandler() *AirGapCommandHandler {
	return &AirGapCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductAirGap),
		AirGapAPI:          coops.NewAirGapAPI("gocean"),
	}
}

// HandleCommand .
func (atch *AirGapCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *AirGapCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *AirGapCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.AirGapAPI.GetAirGap(req)
	if err != nil {
		return fmt.Errorf("could not load air gap data for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.AirGapAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *AirGapCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.AirGapAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve air gap data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
