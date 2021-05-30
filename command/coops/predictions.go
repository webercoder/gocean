package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// PredictionsCommandHandler handles predictions commands.
type PredictionsCommandHandler struct {
	BaseCommandHandler
	predAPI *coops.PredictionsAPI
}

// NewPredictionsCommandHandler creates a new PredictionsCommandHandler.
func NewPredictionsCommandHandler() *PredictionsCommandHandler {
	return &PredictionsCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductPredictions),
		predAPI:            coops.NewPredictionsAPI("gocean"),
	}
}

// HandleCommand processes the predictions command.
func (pch *PredictionsCommandHandler) HandleCommand() error {
	pch.ParseFlags()
	pch.validate()

	req := coops.NewClientRequest(pch.GetRequestOptions()...)

	if pch.clientConfig.Format == ResponseFormatPrettyPrint {
		return pch.handlePrettyPrint(req)
	}

	return pch.handleRawPrint(req)
}

func (pch *PredictionsCommandHandler) validate() {
	if pch.clientConfig.Station == "" {
		pch.Usage(errors.New("station is required"))
	}

	if pch.clientConfig.BeginDate == "" && pch.clientConfig.EndDate == "" {
		pch.clientConfig.BeginDate = time.Now().Format(coops.APIDateFormat)
	}
}

func (pch *PredictionsCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := pch.predAPI.GetPredictions(req)
	if err != nil {
		return fmt.Errorf("could not load predictions for station: %v", err)
	}

	if pch.clientConfig.Count > 0 {
		results = results[:pch.clientConfig.Count]
	}

	pch.predAPI.PrintTabDelimited(pch.clientConfig.Station, results)
	return nil
}

func (pch *PredictionsCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := pch.predAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve predictions data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
