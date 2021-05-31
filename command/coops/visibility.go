package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// VisibilityCommandHandler handles visibility commands.
type VisibilityCommandHandler struct {
	BaseCommandHandler
	VisibilityAPI *coops.VisibilityAPI
}

// NewVisibilityCommandHandler creates a new VisibilityCommandHandler.
func NewVisibilityCommandHandler() *VisibilityCommandHandler {
	return &VisibilityCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductVisibility),
		VisibilityAPI:      coops.NewVisibilityAPI("gocean"),
	}
}

// HandleCommand .
func (atch *VisibilityCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *VisibilityCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *VisibilityCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.VisibilityAPI.GetVisibility(req)
	if err != nil {
		return fmt.Errorf("could not load visibility data for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.VisibilityAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *VisibilityCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.VisibilityAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve visibility data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
