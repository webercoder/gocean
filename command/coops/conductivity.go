package coops

import (
	"errors"
	"fmt"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// ConductivityCommandHandler handles conductivity commands.
type ConductivityCommandHandler struct {
	BaseCommandHandler
	ConductivityAPI *coops.ConductivityAPI
}

// NewConductivityCommandHandler creates a new ConductivityCommandHandler.
func NewConductivityCommandHandler() *ConductivityCommandHandler {
	return &ConductivityCommandHandler{
		BaseCommandHandler: *NewBaseCommandHandler(coops.ProductConductivity),
		ConductivityAPI:    coops.NewConductivityAPI("gocean"),
	}
}

// HandleCommand .
func (atch *ConductivityCommandHandler) HandleCommand() error {
	atch.ParseFlags()
	atch.validate()

	req := coops.NewClientRequest(atch.GetRequestOptions()...)

	if atch.clientConfig.Format == ResponseFormatPrettyPrint {
		return atch.handlePrettyPrint(req)
	}

	return atch.handleRawPrint(req)
}

func (atch *ConductivityCommandHandler) validate() {
	if atch.clientConfig.Station == "" {
		atch.Usage(errors.New("station is required"))
	}

	if atch.clientConfig.BeginDate == "" && atch.clientConfig.EndDate == "" {
		atch.clientConfig.BeginDate = time.Now().Add(-1 * 24 * time.Hour).Format(coops.APIDateFormat)
	}
}

func (atch *ConductivityCommandHandler) handlePrettyPrint(req *coops.ClientRequest) error {
	results, err := atch.ConductivityAPI.GetConductivity(req)
	if err != nil {
		return fmt.Errorf("could not load conductivity data for station: %v", err)
	}

	if atch.clientConfig.Count > 0 {
		results = results[:atch.clientConfig.Count]
	}

	atch.ConductivityAPI.PrintTabDelimited(atch.clientConfig.Station, results)
	return nil
}

func (atch *ConductivityCommandHandler) handleRawPrint(req *coops.ClientRequest) error {
	resp, err := atch.ConductivityAPI.Client.Get(req)
	if err != nil {
		return fmt.Errorf("could not retrieve conductivity data: %v", err)
	}

	fmt.Println(string(resp))
	return nil
}
