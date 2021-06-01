package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// WindCommandHandler handles wind commands.
type WindCommandHandler struct {
	BaseCommandHandler
	WindAPI *coops.WindAPI
}

// NewWindCommandHandler creates a new WindCommandHandler.
func NewWindCommandHandler() *WindCommandHandler {
	return &WindCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductWind),
		WindAPI:            coops.NewWindAPI("gocean"),
	}
}

// HandleCommand .
func (atch *WindCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *WindCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *WindCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.WindAPI.GetWind(req)
	if err != nil {
		return fmt.Errorf("could not load wind data for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.WindAPI.PrintTabDelimited(req, results)
	return nil
}

func (atch *WindCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.WindAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve wind data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
