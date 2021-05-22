package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/webercoder/gocean/src/coops"
)

// PredictionsCommandHandler handles predictions commands.
type PredictionsCommandHandler struct {
	clientConfig *CoopsClientConfig
	flagSet      *flag.FlagSet
	predAPI      *coops.PredictionsAPI
}

// NewPredictionsCommandHandler creates a new Tides and Currents CommandHandler
func NewPredictionsCommandHandler() *PredictionsCommandHandler {
	clientConfig := NewCoopsClientConfig()
	return &PredictionsCommandHandler{
		clientConfig: clientConfig,
		flagSet:      clientConfig.GetFlagSet(coops.ProductPredictions.String(), flag.ExitOnError),
		predAPI:      coops.NewPredictionsAPI("gocean"),
	}
}

// HandleCommand processes the predictions command.
func (pch *PredictionsCommandHandler) HandleCommand() error {
	pch.validate()

	req := coops.NewClientRequest(pch.getRequestOptions()...)

	if pch.clientConfig.Format == ResponseFormatPrettyPrint {
		return pch.handlePrettyPrint(req)
	}

	return pch.handleRawPrint(req)

}

func (pch *PredictionsCommandHandler) validate() {
	if err := pch.flagSet.Parse(os.Args[3:]); err != nil {
		pch.Usage(errors.New("unable to parse command-line options"))
	}

	if pch.clientConfig.Station == "" {
		pch.Usage(errors.New("station is required"))
	}

	if pch.clientConfig.BeginDate == "" && pch.clientConfig.EndDate == "" {
		pch.clientConfig.BeginDate = time.Now().Format(coops.APIDateFormat)
	}
}

func (pch *PredictionsCommandHandler) getRequestOptions() []coops.ClientRequestOption {
	reqOptions, err := pch.clientConfig.ToRequestOptions()
	if err != nil {
		pch.Usage(err)
	}

	return append(reqOptions, coops.WithProduct(coops.ProductPredictions))
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

// Usage prints how to use this command.
func (pch *PredictionsCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	pch.flagSet.Usage()
	os.Exit(1)
}
