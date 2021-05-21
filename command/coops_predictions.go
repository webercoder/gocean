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
		flagSet:      clientConfig.GetFlagSet("predictions", flag.ExitOnError),
		predAPI:      coops.NewPredictionsAPI("gocean"),
	}
}

// HandleCommand processes the predictions command.
func (pch *PredictionsCommandHandler) HandleCommand() error {
	if err := pch.flagSet.Parse(os.Args[3:]); err != nil {
		pch.Usage(errors.New("unable to parse command-line options"))
	}

	if pch.clientConfig.Station == "" {
		pch.Usage(errors.New("station is required"))
	}

	if pch.clientConfig.BeginDate == "" && pch.clientConfig.EndDate == "" {
		pch.clientConfig.BeginDate = time.Now().Format(coops.APIDateFormat)
	}

	reqOptions, err := pch.clientConfig.ToRequestOptions()
	if err != nil {
		pch.Usage(err)
	}
	req := coops.NewClientRequest(
		append(
			reqOptions,
			coops.WithFormat(coops.ResponseFormatJSON),
			coops.WithProduct(coops.ProductPredictions),
		)...,
	)

	results, err := pch.predAPI.Retrieve(req)
	if err != nil {
		return fmt.Errorf("could not load predictions for station: %v", err)
	}
	if pch.clientConfig.Count > 0 {
		results = results[:pch.clientConfig.Count]
	}
	pch.predAPI.PrintTabDelimited(pch.clientConfig.Station, results)
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
