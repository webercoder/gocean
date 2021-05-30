package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// WaterTemperatureCommandHandler handles water levels commands.
type WaterTemperatureCommandHandler struct {
	BaseCommandHandler
	WaterTemperatureAPI *coops.WaterTemperatureAPI
}

// NewWaterTemperatureCommandHandler creates a new WaterTemperatureCommandHandler.
func NewWaterTemperatureCommandHandler() *WaterTemperatureCommandHandler {
	return &WaterTemperatureCommandHandler{
		BaseCommandHandler:  *NewBaseCommandHandler(coops.ProductWaterTemperature),
		WaterTemperatureAPI: coops.NewWaterTemperatureAPI("gocean"),
	}
}

// HandleCommand .
func (atch *WaterTemperatureCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *WaterTemperatureCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
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
