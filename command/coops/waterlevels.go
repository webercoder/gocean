package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// WaterLevelsCommandHandler handles water levels commands.
type WaterLevelsCommandHandler struct {
	BaseCommandHandler
	waterLevelAPI *coops.WaterLevelAPI
}

// NewWaterLevelsCommandHandler creates a new WaterLevelsCommandHandler.
func NewWaterLevelsCommandHandler() *WaterLevelsCommandHandler {
	return &WaterLevelsCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductWaterLevel),
		waterLevelAPI:      coops.NewWaterLevelAPI("gocean"),
	}
}

// HandleCommand .
func (wlch *WaterLevelsCommandHandler) HandleCommand() error {
	wlch.ParseFlags()
	wlch.validate()

	req := coops.NewClientRequest(wlch.GetRequestOptions()...)

	if wlch.clientConfig.Format == ResponseFormatPrettyPrint {
		return wlch.handlePrettyPrint(req)
	}

	return wlch.handleRawPrint(req)
}

func (wlch *WaterLevelsCommandHandler) validate() {
	if wlch.clientConfig.Station == "" {
		wlch.Usage(errors.New("station is required"))
	}

	if wlch.clientConfig.BeginDate == "" && wlch.clientConfig.EndDate == "" {
		wlch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (wlch *WaterLevelsCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := wlch.waterLevelAPI.GetWaterLevels(req)
	if err != nil {
		return fmt.Errorf("could not load water levels for station: %v", err)
	}

	if wlch.clientConfig.Count > 0 {
		results = results[:wlch.clientConfig.Count]
	}

	wlch.waterLevelAPI.PrintTabDelimited(req, results)
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
